package tests

import (
	"encoding/json"
	"github.com/dmitry-udod/codes_go/cmd"
	"github.com/matryer/is"
	"log"
	"os"
	"strings"
	"testing"
)

func TestFopXmlStringDecode(t *testing.T) {
	recordStr := `<RECORD><FIO>САЄНКО ОЛЬГА СЕРГІЇВНА</FIO><ADDRESS>91007, Луганська обл., місто Луганськ, Артемівський район, ВУЛИЦЯ ПРИВОЗНА, будинок 55</ADDRESS><KVED>47.82 Роздрібна Торгівля з лотків і на ринках текстильними виробами, одягом і взуттям</KVED><STAN>припинено</STAN></RECORD>`
	record := cmd.DecodeFopXmlString(recordStr)

	assert := is.New(t)
	assert.Equal("Саєнко Ольга Сергіївна", record.FullName)
	assert.Equal("91007, Луганська Обл., Місто Луганськ, Артемівський Район, Вулиця Привозна, Будинок 55", record.Address)
	assert.Equal("47.82 роздрібна торгівля з лотків і на ринках текстильними виробами, одягом і взуттям", record.Activity)
	assert.Equal("припинено", record.Status)
	assert.Equal(TEST_DOCUMENT_ID, record.GenerateId())

	data, err := json.Marshal(record)
	assert.True(err == nil)
	assert.Equal(`{"full_name":"Саєнко Ольга Сергіївна","address":"91007, Луганська Обл., Місто Луганськ, Артемівський Район, Вулиця Привозна, Будинок 55","activity":"47.82 роздрібна торгівля з лотків і на ринках текстильними виробами, одягом і взуттям","status":"припинено"}`, string(data))
}

func TestTerroristXmlDecode(t *testing.T) {
	file, err := os.Open("./mocks/terrorists.xml")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	list, _ := cmd.ParseTerroristXmlFile(file)

	assert := is.New(t)
	assert.Equal(1133, list.Length)
	assert.Equal("364", list.Version)
	assert.Equal("19.07.2019", list.LastUpdated)
	assert.Equal(2, len(list.Terrorists))
	assert.Equal(3, list.Terrorists[0].Number)
	assert.Equal(4, list.Terrorists[1].Number)
	assert.Equal("13.10.2010", list.Terrorists[0].AddedAt)
	assert.Equal("20.05.2011", list.Terrorists[1].AddedAt)
	assert.Equal("Офіційний веб-сайт ООН", list.Terrorists[0].Source)
	assert.Equal("Резолюції РБ ООН 1267(1999), 1904 (2009)", list.Terrorists[1].Source)
	assert.Equal(6, len(list.Terrorists[0].AkaNames))
	assert.Equal(9, len(list.Terrorists[1].AkaNames))
	assert.Equal("Mohammad", list.Terrorists[0].AkaNames[0].LastName)
	assert.Equal("Hamdi", list.Terrorists[0].AkaNames[0].FirstName)
	assert.Equal("19 Nov. 1971", list.Terrorists[0].BirthDay)
	assert.Equal("18 Dec. 1969", list.Terrorists[1].BirthDay)
	assert.Equal("Medina, Saudi Arabia", list.Terrorists[0].BirthPlaces[0])
	assert.Equal("Asima-Tunis, Tunisia", list.Terrorists[1].BirthPlaces[0])
	assert.Equal("born 25 May 1968 in Naples, Italy", list.Terrorists[1].BirthPlaces[1])
	assert.Equal("Yemeni", list.Terrorists[0].Nationalities[0])
	assert.Equal("Tunisian", list.Terrorists[1].Nationalities[0])
	assert.True(strings.Contains(list.Terrorists[0].Comments, "Yemeni passport number 541939 issued in Al-Hudaydah"))
	assert.True(strings.Contains(list.Terrorists[1].Comments, "Father’s name is Mahmoud ben Sasi"))
}
