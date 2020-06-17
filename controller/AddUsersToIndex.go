package controller

import (
	"sync"
	"log"

	model "github.com/nvhai245/goelastic/model"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

// AddUsersToIndex concurrently
func AddUsersToIndex(es *elasticsearch.Client, users []model.NewUser) (result bool) {
	var wg sync.WaitGroup
	for i, user := range users {
		wg.Add(1)

		res, err := es.Index("user", esutil.NewJSONReader(user), es.Index.WithRefresh("true"), es.Index.WithPretty())
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
			return false
		}
		if res.IsError() {
			log.Fatalf("Can not index new user, response: %s", res)
		}
		go func(i int, user model.NewUser) {
			defer wg.Done()
		}(i, user)
	}
	return true
}
