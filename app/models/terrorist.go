package models

//const INDEX_TERRORISTS = "terrorists"

type Terrorist struct {
	Number   uint      `xml:"number-entry" json:"number_in_list"`
	AddedAt  string    `xml:"date-entry" json:"added_at"`
	Source   string    `xml:"program-entry" json:"source"`
	AkaNames []AkaName `xml:"aka-list"`
}

type AkaName struct {
}
