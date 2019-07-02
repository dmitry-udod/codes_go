package tests

import (
	"encoding/json"
	"github.com/dmitry-udod/codes_go/models"
	"github.com/dmitry-udod/codes_go/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

const TEST_INDEX = "tests"
const TEST_DOCUMENT_ID = "1568827652"

func TestSaveAndSearchData(t *testing.T) {
	checkConnectionToEsServer(t)

	m := make(map[string]string)
	record := new(models.Record)
	record.FullName = "Саєнко Ольга Сергіївна"
	record.Address = "91007, Луганська Обл., Місто Луганськ, Артемівський Район, Вулиця Привозна, Будинок 55"
	jsonString, _ := json.Marshal(record)
	m[record.GenerateId()] = string(jsonString)

	resp := services.SaveDataToEs(TEST_INDEX, m)
	assert.Equal(t, 201, resp.Items[0].Index.Status)
	assert.Equal(t, TEST_DOCUMENT_ID, resp.Items[0].Index.ID)

	entities := services.Search(TEST_INDEX, TEST_DOCUMENT_ID)
	search := new(models.Record)
	search.ParseFromSearch(entities[0])

	assert.Equal(t, record.Address, search.Address)
	assert.Equal(t, record.FullName, search.FullName)
}

func checkConnectionToEsServer(t *testing.T) bool {
	isConnected := services.InitElasticSearchClient()
	
	if (isConnected) {
		clearTestIndex()
	} else {
		t.Skip("Elastic Search server not running")
	}
	
	return isConnected
}

func clearTestIndex()  {
	services.DeleteDataFromEs(TEST_INDEX, TEST_DOCUMENT_ID)
}