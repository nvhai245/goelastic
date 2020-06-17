package controller

import (
	_ "fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

// CreateIndex with built-in analyzer
func CreateIndex(es *elasticsearch.Client, index string, analyzer string) (result bool) {
	// Check if index already exist
	res, err := es.Indices.Exists([]string{index}, es.Indices.Exists.WithPretty())
	if err != nil {
		log.Printf("Error checking indices status: %s", err)
		return false
	}
	if res.StatusCode == 200 {
		log.Printf(`Index "%s" already exist`, index)
		return false
	}

	// Check if analyzer is valid
	builtInAnalyzers := []string{"standard", "simple", "whitespace", "stop", "keyword", "pattern", "fingerprint"}
	contains := false
	for _, builtInAnalyzer := range builtInAnalyzers {
		if builtInAnalyzer == analyzer {
			contains = true
			break
		}
	}
	if !contains {
		log.Printf(`"%s" is not a buit-in analyzer`, analyzer)
		return false
	}
	settingString := `{
			"settings": {
				"analysis": {
				  "analyzer": {
					"my_analyzer": {
					  "type": "custom",
					  "tokenizer": "keyword",
					  "filter": [
						"lowercase",
						"asciifolding"
					  ]
					}
				  }
				}
			  },
			  "mappings": {
				  "properties": {
					"first_name": {
					  "type": "text",
					  "analyzer": "my_analyzer"
				    }
				}
			  }
		}`
	indexSettings := settingString
	res, err = es.Indices.Create(index, es.Indices.Create.WithBody(strings.NewReader(indexSettings)))
	if err != nil {
		log.Printf(`Error create index "%s": %s`, index, err)
		return false
	}
	if res.IsError() {
		log.Printf(`Error create index, response: %s`, res)
		return false
	}
	log.Println(res)
	return true
}
