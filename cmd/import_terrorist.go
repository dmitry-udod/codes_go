package cmd

import (
	"encoding/xml"
	"fmt"
	"github.com/dmitry-udod/codes_go/app/models"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"os"
)

func ImportTerrorist(file *os.File) TerroristList {
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel
	content, _ := ioutil.ReadAll(file)
	var list TerroristList
	xml.Unmarshal(content, &list)

	if len(list.LastUpdated) < 8 {
		log.Fatalf("Cant set correct update data for: %s", list.LastUpdated)
	}

	list.LastUpdated = convertDate(list.LastUpdated)

	for index, t := range list.Terrorists {
		list.Terrorists[index].AddedAt = convertDate(t.AddedAt)
	}

	fmt.Printf("%v", list)
	return list
}

func convertDate(s string) string {
	return s[6:8] + "." + s[4:6] + "." + s[:4]
}

type TerroristList struct {
	XMLName     xml.Name           `xml:"list-terror"`
	Length      uint               `xml:"count-record"`
	Version     string             `xml:"ver-list"`
	LastUpdated string             `xml:"date-ver-list"`
	Terrorists  []models.Terrorist `xml:"acount-list"`
}
