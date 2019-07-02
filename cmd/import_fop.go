package cmd

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	. "github.com/dmitry-udod/codes_go/logger"
	. "github.com/dmitry-udod/codes_go/models"
	es "github.com/dmitry-udod/codes_go/services"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"log"
	"os"
	"strings"
)

func ImportFop(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		msg := fmt.Sprintf("File %s NOT found", filePath)
		fmt.Println(msg)
		Log.Fatal(msg);
	}

	Log.Info("Start processing file: " + filePath);
	file, err := os.Open(filePath)
	if err != nil {
		Log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linesCount := 0
	bulk := 0

	pack := make(map[string]string, 0)

	for scanner.Scan() {
		linesCount++

		if (linesCount == 1) {
			continue
		}

		bulk++

		utfString, _, err := transform.String(charmap.Windows1251.NewDecoder(), scanner.Text())

		if err != nil {
			Log.Fatal(fmt.Sprintf("Error while string decode %s. On Line: %d", err.Error(), ))
		}

		fmt.Println(fmt.Sprintf("Process line number: %d", linesCount))

		record := DecodeFopXmlString(utfString)

		jsonData, err := json.Marshal(record)

		if (err != nil) {
			Log.Error("Cant marshal record to json", record)
		} else {
			pack[record.GenerateId()] = string(jsonData)
		}

		if bulk > 10000 {
			fmt.Printf("[ELASTIC] Save data bulk")
			es.SaveDataToEs("fops", pack)
			bulk = 0
			pack = make(map[string]string, 0)
		}
	}

	Log.Info("Finish processing file");

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func DecodeFopXmlString(str string) Record {
	var record Record

	err := xml.Unmarshal([]byte(str), &record)

	if err != nil {
		msg := fmt.Sprintf("Cant decode XML record: %s", str)
		fmt.Println(msg)
		Log.Fatal(msg)
	}

	record.FullName = strings.Title(strings.ToLower(record.FullName))
	record.Address = strings.Title(strings.ToLower(record.Address))
	record.Activity = strings.ToLower(record.Activity)

	return record
}
