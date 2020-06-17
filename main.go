package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang/glog"
	controller "github.com/nvhai245/goelastic/controller"
	model "github.com/nvhai245/goelastic/model"

	"github.com/elastic/go-elasticsearch/v7"
	_ "github.com/elastic/go-elasticsearch/v7/esapi"
	_ "github.com/elastic/go-elasticsearch/v7/esutil"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func main() {
	log.SetFlags(0)
	flag.Parse()
	var (
		err error
	)

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Create user index
	result := controller.CreateIndex(es, "user", "standard")
	if result == false {
		log.Println("Create index failed")
	}

	// Indexing new user
	result = controller.AddUsersToIndex(es, []model.NewUser{model.User1, model.User2, model.User3, model.User4})
	if !result {
		log.Println("Indexing users failed")
	}

	foundUsers := controller.SearchUsername(es, "user", "Tim")
	if len(foundUsers.ListUserID) == 0 {
		log.Println("Search for users failed")
	}
	glog.Flush()
}
