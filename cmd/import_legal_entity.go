package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/dmitry-udod/codes_go/app/models"
	. "github.com/dmitry-udod/codes_go/logger"
	"golang.org/x/net/html/charset"
	"os"
	"strings"
	es "github.com/dmitry-udod/codes_go/app/services"
)

func ImportLegalEntity(file *os.File) {
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel
	bulk := 0
	linesCount := 0
	pack := make(map[string]string, 0)

	for {
		// Read tokens from the XML document in a stream.
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

				var record models.LegalEntity
				err := decoder.DecodeElement(&record, &se)

				if err != nil {
					Log.Error(err)
					continue
				}

				bulk++
				record.FullName = strings.Title(strings.ToLower(record.FullName))
				record.ShortName = strings.Title(strings.ToLower(record.ShortName))
				record.Address = strings.Title(strings.ToLower(record.Address))
				record.Director = strings.Title(strings.ToLower(record.Director))
				record.Activity = strings.ToLower(record.Activity)

				if len(record.Founders) > 0 {
					for _, f := range record.Founders {
						f.Name = strings.Title(strings.ToLower(f.Name))
					}
				}

				if record.FullName != "" {
					jsonData, err := json.Marshal(record)

					if err == nil {
						pack[record.Code] = string(jsonData)
					}
				}

				if bulk > 10000 {
					fmt.Printf("[ELASTIC] Save data bulk")
					es.SaveDataToEs(models.INDEX_LEGAL_ENTITY, pack)
					bulk = 0
					pack = make(map[string]string, 0)
				}
			}
		}
	}

}