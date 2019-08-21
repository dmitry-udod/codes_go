package tests

import (
	"encoding/json"
	"github.com/dmitry-udod/codes_go/cmd"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestFopXmlStringDecode(t *testing.T) {

	recordStr := `<RECORD><FIO>САЄНКО ОЛЬГА СЕРГІЇВНА</FIO><ADDRESS>91007, Луганська обл., місто Луганськ, Артемівський район, ВУЛИЦЯ ПРИВОЗНА, будинок 55</ADDRESS><KVED>47.82 Роздрібна Торгівля з лотків і на ринках текстильними виробами, одягом і взуттям</KVED><STAN>припинено</STAN></RECORD>`

	record := cmd.DecodeFopXmlString(recordStr)

	assert.Equal(t, "Саєнко Ольга Сергіївна", record.FullName)
	assert.Equal(t, "91007, Луганська Обл., Місто Луганськ, Артемівський Район, Вулиця Привозна, Будинок 55", record.Address)
	assert.Equal(t, "47.82 роздрібна торгівля з лотків і на ринках текстильними виробами, одягом і взуттям", record.Activity)
	assert.Equal(t, "припинено", record.Status)
	assert.Equal(t, TEST_DOCUMENT_ID, record.GenerateId())

	data, err := json.Marshal(record)
	assert.Nil(t, err)
	assert.Equal(t, `{"full_name":"Саєнко Ольга Сергіївна","address":"91007, Луганська Обл., Місто Луганськ, Артемівський Район, Вулиця Привозна, Будинок 55","activity":"47.82 роздрібна торгівля з лотків і на ринках текстильними виробами, одягом і взуттям","status":"припинено"}`, string(data))
}

func TestTerroristXmlDecode(t *testing.T) {
	file, err := os.Open("./mocks/terrorists.xml")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	list := cmd.ImportTerrorist(file)

	assert.Equal(t, uint(1133), list.Length)
	assert.Equal(t, "364", list.Version)
	assert.Equal(t, "19.07.2019", list.LastUpdated)
	assert.Equal(t, 2, len(list.Terrorists))
	assert.Equal(t, uint(3), list.Terrorists[0].Number)
	assert.Equal(t, uint(4), list.Terrorists[1].Number)
	assert.Equal(t, "13.10.2010", list.Terrorists[0].AddedAt)
	assert.Equal(t, "20.05.2011", list.Terrorists[1].AddedAt)
	assert.Equal(t, "Офіційний веб-сайт ООН", list.Terrorists[0].Source)
	assert.Equal(t, "Резолюції РБ ООН 1267(1999), 1904 (2009)", list.Terrorists[1].Source)
	assert.Equal(t, 6, len(list.Terrorists[0].AkaNames))
	assert.Equal(t, 9, len(list.Terrorists[1].AkaNames))
	assert.Equal(t, "Mohammad", list.Terrorists[0].AkaNames[0].LastName)
	assert.Equal(t, "Hamdi", list.Terrorists[0].AkaNames[0].FirstName)
	assert.Equal(t, "19 Nov. 1971", list.Terrorists[0].BirthDay)
	assert.Equal(t, "18 Dec. 1969", list.Terrorists[1].BirthDay)
	assert.Equal(t, "Medina, Saudi Arabia", list.Terrorists[0].BirthPlaces[0])
	assert.Equal(t, "Asima-Tunis, Tunisia", list.Terrorists[1].BirthPlaces[0])
	assert.Equal(t, "born 25 May 1968 in Naples, Italy", list.Terrorists[1].BirthPlaces[1])
	assert.Equal(t, "Yemeni", list.Terrorists[0].Nationalities[0])
	assert.Equal(t, "Tunisian", list.Terrorists[1].Nationalities[0])
	assert.Contains(t, list.Terrorists[0].Comments, "Yemeni passport number 541939 issued in Al-Hudaydah")
	assert.Contains(t, list.Terrorists[1].Comments, "Father’s name is Mahmoud ben Sasi")
}
