package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	. "github.com/dmitry-udod/codes_go/app/models"
	es "github.com/dmitry-udod/codes_go/app/services"
	. "github.com/dmitry-udod/codes_go/logger"
	"golang.org/x/net/html/charset"
	"os"
	"strings"
)

func ImportFop(file *os.File) {
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel
	linesCount := 0
	bulk := 0

	pack := make(map[string]string, 0)

	for {
		t, err := decoder.Token()
		if t == nil {
			break
		}

		if err != nil {
			Log.Fatalf("cant parse file: %s", err.Error())
			return
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "RECORD" {
				linesCount++
				fmt.Println(fmt.Sprintf("Process line number: %d", linesCount))

				var record Record
				err := decoder.DecodeElement(&record, &se)

				if err != nil {
					Log.Error(err)
					continue
				}

				bulk++

				record = decorateFopRecord(record)

				if record.FullName != "" {
					jsonData, err := json.Marshal(record)

					if err == nil {
						pack[record.GenerateId()] = string(jsonData)
					}
				}

				if bulk > 10000 {
					fmt.Printf("[ELASTIC] Save data bulk")
					es.SaveDataToEs(INDEX_FOP, pack)
					bulk = 0
					pack = make(map[string]string, 0)
				}
			}
		}
	}

	fmt.Printf("[ELASTIC] Save data bulk")
	es.SaveDataToEs(INDEX_FOP, pack)
}

func DecodeFopXmlString(str string) Record {
	var record Record

	err := xml.Unmarshal([]byte(str), &record)

	if err != nil {
		msg := fmt.Sprintf("Cant decode XML record: %s", str)
		fmt.Println(msg)
		Log.Errorf(msg)
	}

	record = decorateFopRecord(record)

	return record
}

func decorateFopRecord(record Record) Record {
	record.FullName = strings.Title(strings.ToLower(record.FullName))
	record.Address = strings.Title(strings.ToLower(record.Address))
	record.Activity = strings.ToLower(record.Activity)

	return record
}
