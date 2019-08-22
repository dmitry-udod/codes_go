package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/dmitry-udod/codes_go/app/models"
	es "github.com/dmitry-udod/codes_go/app/services"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseTerroristXmlFile(file *os.File) (TerroristList, map[string]string) {
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel
	content, _ := ioutil.ReadAll(file)
	var list TerroristList
	xml.Unmarshal(content, &list)

	if len(list.LastUpdated) < 8 {
		log.Fatalf("Cant set correct update data for: %s", list.LastUpdated)
	}

	list.LastUpdated = convertDate(list.LastUpdated)

	pack := make(map[string]string, 0)
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

		jsonData, err := json.Marshal(t)

		if err == nil {
			pack[strconv.Itoa(t.Number)] = string(jsonData)
		}
	}

	return list, pack
}

func ImportTerrorist(file *os.File) {
	_, pack := ParseTerroristXmlFile(file)

	if len(pack) > 0 {
		fmt.Printf("[ELASTIC] Save %s data ", models.INDEX_TERRORISTS)
		es.SaveDataToEs(models.INDEX_TERRORISTS, pack)
	}
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
	Length      int                `xml:"count-record"`
	Version     string             `xml:"ver-list"`
	LastUpdated string             `xml:"date-ver-list"`
	Terrorists  []models.Terrorist `xml:"acount-list"`
}
