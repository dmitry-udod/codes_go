package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dmitry-udod/codes_go/app/models"
	. "github.com/dmitry-udod/codes_go/logger"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strconv"
	"strings"
)

var Es *elasticsearch.Client

func InitElasticSearchClient() bool {
	var r map[string]interface{}
	var err error

	// Initialize a client with the default settings.
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	Es, err = elasticsearch.NewDefaultClient()
	if err != nil {
		Log.Errorf("Error creating the client: %s", err)
		return false
	}

	res, err := Es.Info()
	if err != nil {
		Log.Errorf("Error getting response: %s", err)
		return false
	}

	// Check response status
	if res.IsError() {
		Log.Errorf("Error: %s", res.String())
		return false
	}

	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		Log.Errorf("Error parsing the response body: %s", err)
		return false
	}

	Log.Printf("Successfully connected to Server: %s", r["version"].(map[string]interface{})["number"])
	Log.Println(strings.Repeat("~", 37))

	return true
}

func Search(index string, params map[string]string) ([]interface{}, models.Metadata) {
	id, idExist := params["id"]
	page, _ := strconv.Atoi(params["page"])
	perPage := 10

	var r map[string]interface{}
	var buf bytes.Buffer

	query := map[string]interface{}{}
	if idExist {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"_id": id,
				},
			},
		}
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		Log.Errorf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := Es.Search(
		Es.Search.WithContext(context.Background()),
		Es.Search.WithIndex(index),
		Es.Search.WithBody(&buf),
		Es.Search.WithTrackTotalHits(true),
		Es.Search.WithPretty(),
		Es.Search.WithSize(perPage),
		Es.Search.WithFrom(page*perPage),
	)
	if err != nil {
		Log.Errorf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			Log.Errorf("Error parsing the response body: %s", err)
		} else {
			Log.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		Log.Errorf("Error parsing the response body: %s", err)
	}

	total := int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	Log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		total,
		int(r["took"].(float64)),
	)

	metadata := models.Metadata{
		Total: uint(total),
	}

	return r["hits"].(map[string]interface{})["hits"].([]interface{}), metadata
}

func SaveDataToEs(index string, data map[string]string) *bulkResponse {
	var buf bytes.Buffer
	var blk *bulkResponse
	var raw map[string]interface{}

	for id, entity := range data {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s" } }%s`, id, "\n"))

		entityByte := []byte(entity + "\n")

		buf.Grow(len(meta) + len(entityByte))
		buf.Write(meta)
		buf.Write(entityByte)
	}

	res, err := Es.Bulk(bytes.NewReader(buf.Bytes()), Es.Bulk.WithIndex(index))
	if err != nil {
		Log.Fatalf("Failure indexing batch %s", err)
	}

	if res.IsError() {
		if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
			Log.Fatalf("Failure to to parse response body: %s", err)
		} else {
			Log.Infof("  Error: [%d] %s: %s",
				res.StatusCode,
				raw["error"].(map[string]interface{})["type"],
				raw["error"].(map[string]interface{})["reason"],
			)
		}
	} else {
		if err := json.NewDecoder(res.Body).Decode(&blk); err != nil {
			Log.Fatalf("Failure to to parse response body: %s", err)
		} else {
			for _, d := range blk.Items {
				if d.Index.Status > 201 {
					Log.Info("  Error: [%d]: %s: %s: %s: %s",
						d.Index.Status,
						d.Index.Error.Type,
						d.Index.Error.Reason,
						d.Index.Error.Cause.Type,
						d.Index.Error.Cause.Reason,
					)
				}
			}
		}
	}

	buf.Reset()

	return blk
}

func DeleteDataFromEs(index, id string) bool {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), Es)
	if err != nil {
		Log.Errorf("Error getting response: %s", err)
		return false
	}
	defer res.Body.Close()

	if res.IsError() {
		Log.Errorf("[%s] Error while request to ES: %v", res.Status(), req)
		return false
	}

	Log.Errorf("[%s] Delete document %s from index", res.Status(), id)

	return true
}

func SearchFop(params map[string]string) ([]*models.RecordWithId, models.Metadata) {
	records := make([]*models.RecordWithId, 0)
	metadata := models.Metadata{}

	if ! InitElasticSearchClient() {
		Log.Error("ES server iis not available")
		return records, metadata
	}

	entities, metadata := Search(models.INDEX_FOP, params)

	if len(entities) > 0 {
		for _, entity := range entities {
			record := models.Record{}
			record.ParseFromSearch(entity)
			recordWithId := new(models.RecordWithId)
			recordWithId.Record = record
			recordWithId.Id = record.GenerateId()
			records = append(records, recordWithId)
		}
	}

	return records, metadata
}

type bulkResponse struct {
	Errors bool `json:"errors"`
	Items  []struct {
		Index struct {
			ID     string `json:"_id"`
			Result string `json:"result"`
			Status int    `json:"status"`
			Error  struct {
				Type   string `json:"type"`
				Reason string `json:"reason"`
				Cause  struct {
					Type   string `json:"type"`
					Reason string `json:"reason"`
				} `json:"caused_by"`
			} `json:"error"`
		} `json:"index"`
	} `json:"items"`
}
