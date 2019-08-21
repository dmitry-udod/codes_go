package cmd

import (
	"encoding/xml"
	"github.com/dmitry-udod/codes_go/app/models"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
		list.Terrorists[index].BirthDay = trim(t.BirthDay)
		for akaIndex, akaName := range t.AkaNames {
			list.Terrorists[index].AkaNames[akaIndex].FirstName = convertName(akaName.FirstName)
			list.Terrorists[index].AkaNames[akaIndex].LastName = convertName(akaName.LastName)
		}

		for birthdayIndex, birthdayPlace := range t.BirthPlaces {
			list.Terrorists[index].BirthPlaces[birthdayIndex] = trim(birthdayPlace)
		}
	}

	return list
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func convertDate(s string) string {
	return s[6:8] + "." + s[4:6] + "." + s[:4]
}

func convertName(s string) string {
	return strings.Title(strings.ToLower(s))
}

type TerroristList struct {
	XMLName     xml.Name           `xml:"list-terror"`
	Length      uint               `xml:"count-record"`
	Version     string             `xml:"ver-list"`
	LastUpdated string             `xml:"date-ver-list"`
	Terrorists  []models.Terrorist `xml:"acount-list"`
}
