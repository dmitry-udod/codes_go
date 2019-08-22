package models

const INDEX_TERRORISTS = "terrorists"

type Terrorist struct {
	Number        int       `xml:"number-entry" json:"number_in_list"`
	AddedAt       string    `xml:"date-entry" json:"added_at"`
	Source        string    `xml:"program-entry" json:"source"`
	AkaNames      []AkaName `xml:"aka-list" json:"known_names"`
	BirthDay      string    `xml:"date-of-birth-list" json:"birth_day"`
	BirthPlaces   []string  `xml:"place-of-birth-list" json:"birth_places"`
	Nationalities []string  `xml:"nationality-list" json:"nationalities"`
	Comments      string    `xml:"comments" json:"comments"`
}

type AkaName struct {
	LastName       string `xml:"aka-name1" json:"last_name"`
	FirstName      string `xml:"aka-name2" json:"first_name"`
	MiddleName     string `xml:"aka-name3" json:"middle_name"`
	AdditionalName string `xml:"aka-name4" json:"additional_name"`
}
