package controller

import (
	"encoding/json"
	"log"

	model "github.com/nvhai245/goelastic/model"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

// SearchUsername func
func SearchUsername(es *elasticsearch.Client, index string, searchQuery string) (listUserID model.ListUserID) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase_prefix": map[string]interface{}{
				"first_name": map[string]interface{}{
					"query":    "Tim",
				},
			},
		},
	}

	res, err := es.Search(
		es.Search.WithIndex(index),
		es.Search.WithBody(esutil.NewJSONReader(&query)),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return model.ListUserID{}
	}
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	defer res.Body.Close()

	// Deserialize the response into a struct.
	r := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	jsonString, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	s := model.SearchResult{}
	if err := json.Unmarshal(jsonString, &s); err != nil {
        log.Println(err)
	}
	log.Println(s)
	return model.ListUserID{}
}
