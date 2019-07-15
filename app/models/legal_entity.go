package models

import (
	"encoding/json"
	"encoding/xml"
)

const INDEX_LEGAL_ENTITY = "legal_entities"

type Founder struct {
	Name string `xml:"FOUNDER" json:"name"`
}

type LegalEntity struct {
	XMLName   xml.Name   `xml:"RECORD" json:"-"`
	FullName  string     `xml:"NAME" json:"full_name"`
	ShortName string     `xml:"SHORT_NAME" json:"short_name"`
	Code      string     `xml:"EDRPOU" json:"code"`
	Address   string     `xml:"ADDRESS" json:"address"`
	Director  string     `xml:"BOSS" json:"director"`
	Activity  string     `xml:"KVED" json:"activity"`
	Status    string     `xml:"STAN" json:"status"`
	Founders  []*Founder `xml:"FOUNDERS" json:"founders"`
}

func (r *LegalEntity) ParseFromSearch(search interface{}) {
	source, _ := json.Marshal(search.(map[string]interface{})["_source"])
	json.Unmarshal(source, &r)
}
