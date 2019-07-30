package cmd

import (
	"encoding/xml"
	"fmt"
	"github.com/dmitry-udod/codes_go/app/models"
	. "github.com/dmitry-udod/codes_go/logger"
	"golang.org/x/net/html/charset"
	"os"
)

const DOMAIN = "https://da.org.ua"

func GenerateSiteMap(file *os.File) {
	url := DOMAIN + "/#/legal-entities/details/"
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel
	bulk := 0
	linesCount := 0

	f, err := os.Create("public/sitemap.txt")

	if err != nil {
		Log.Fatal(err.Error())
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("%s%s\n", DOMAIN, "/#/fop"))
	f.WriteString(fmt.Sprintf("%s%s\n", DOMAIN, "/#/legal-entities"))

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

				f.WriteString(fmt.Sprintf("%s%s\n", url, record.Code))

			}
		}

	}

}