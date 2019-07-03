package models

import (
	"encoding/json"
	"encoding/xml"
	"hash/fnv"
	"strconv"
)

const INDEX_FOP = "fops"

type Record struct {
	XMLName  xml.Name `xml:"RECORD" json:"-"`
	FullName string   `xml:"FIO" json:"full_name"`
	Address  string   `xml:"ADDRESS" json:"address"`
	Activity string   `xml:"KVED" json:"activity"`
	Status   string   `xml:"STAN" json:"status"`
}

func (r *Record) GenerateId() string {
	text := r.FullName + r.Address
	algorithm := fnv.New32a()
	algorithm.Write([]byte(text))
	return strconv.FormatUint(uint64(algorithm.Sum32()), 10)
}

func (r *Record) ParseFromSearch(search interface{}) {
	source, _ := json.Marshal(search.(map[string]interface{})["_source"])
	json.Unmarshal(source, &r)
}
