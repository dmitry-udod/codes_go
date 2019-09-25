package tests

import (
	"encoding/json"
	"github.com/dmitry-udod/codes_go/app/controllers"
	"github.com/dmitry-udod/codes_go/app/models"
	"github.com/dmitry-udod/codes_go/app/services"
	"github.com/matryer/is"
	"testing"
)

const TEST_INDEX = "tests"
const TEST_DOCUMENT_ID = "1568827652"
const TEST_FOP_ID = "4237321928"
const TEST_FOP_NAME = "Столяров Руслан Миколайович"

func TestSaveAndSearchData(t *testing.T) {
	checkConnectionToEsServer(t)

	m := make(map[string]string)
	record := new(models.Record)
	record.FullName = "Саєнко Ольга Сергіївна"
	record.Address = "91007, Луганська Обл., Місто Луганськ, Артемівський Район, Вулиця Привозна, Будинок 55"
	jsonString, _ := json.Marshal(record)
	m[record.GenerateId()] = string(jsonString)

	assert := is.New(t)
	resp := services.SaveDataToEs(TEST_INDEX, m)
	assert.Equal(201, resp.Items[0].Index.Status)
	assert.Equal(TEST_DOCUMENT_ID, resp.Items[0].Index.ID)

	params := controllers.Params()
	params["id"] = TEST_DOCUMENT_ID

	entities, metadata := services.Search(TEST_INDEX, params)
	search := new(models.Record)
	search.ParseFromSearch(entities[0])

	assert.Equal(record.Address, search.Address)
	assert.Equal(record.FullName, search.FullName)
	assert.Equal(uint(1), metadata.Total)
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

func clearTestIndex() {
	services.DeleteDataFromEs(TEST_INDEX, TEST_DOCUMENT_ID)
}