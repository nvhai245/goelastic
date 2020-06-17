package model

import (
	"encoding/json"
	_ "log"
	"github.com/google/uuid"
)

// NewUser object
type NewUser struct {
	ID                   int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string            `protobuf:"bytes,11,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string            `protobuf:"bytes,12,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FullName             string            `protobuf:"bytes,13,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Username             string            `protobuf:"bytes,14,opt,name=username,proto3" json:"username,omitempty"`
	Bio                  string            `protobuf:"bytes,15,opt,name=bio,proto3" json:"bio,omitempty"`
}

// ListUserID to be returned
type ListUserID struct {
	ListUserID []int32
}


// Response represents a boolean response sent back by the search egine
type Response struct {
	Acknowledged bool
	Error        string
	Status       int
}

// Settings represents the mapping structure of one or several indices
type Settings struct {
	Shards  map[string]interface{} `json:"_shards"`
	Indices map[string]interface{} `json:"indices"`
}

// Status represents the status of the search engine
type Status struct {
	TagLine string
	Version struct {
		Number         string
		BuildHash      string `json:"build_hash"`
		BuildTimestamp string `json:"build_timestamp"`
		BuildSnapshot  bool   `json:"build_snapshot"`
		LuceneVersion  string `json:"lucene_version"`
	}
	Name   string
	Status int
	Ok     bool
}

// InsertDocument represents the result of the insert operation of a document
type InsertDocument struct {
	Created bool   `json:"created"`
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
}

// Document represents a document
type Document struct {
	Index   string          `json:"_index"`
	Type    string          `json:"_type"`
	ID      string          `json:"_id"`
	Version int             `json:"_version"`
	Found   bool            `json:"found"`
	Source  json.RawMessage `json:"_source"`
}

// Bulk represents the result of the Bulk operation
type Bulk struct {
	Took   uint64 `json:"took"`
	Errors bool   `json:"errors"`
	Items  []struct {
		Create struct {
			Index  string `json:"_index"`
			Type   string `json:"_type"`
			ID     string `json:"_id"`
			Status int    `json:"status"`
			Error  string `json:"error"`
		} `json:"create"`
		Index struct {
			Index   string `json:"_index"`
			Type    string `json:"_type"`
			ID      string `json:"_id"`
			Version int    `json:"_version"`
			Status  int    `json:"status"`
			Error   string `json:"error"`
		} `json:"index"`
	} `json:"items"`
}

// SearchResult represents the result of the search operation
type SearchResult struct {
	Took     uint64 `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits         ResultHits      `json:"hits"`
	Aggregations json.RawMessage `json:"aggregations"`
}

// ResultHits represents the result of the search hits
type ResultHits struct {
	Total    int     `json:"total"`
	MaxScore float32 `json:"max_score"`
	Hits     []struct {
		Index     string              `json:"_index"`
		Type      string              `json:"_type"`
		ID        string              `json:"_id"`
		Score     float32             `json:"_score"`
		Source    json.RawMessage     `json:"_source"`
		Highlight map[string][]string `json:"highlight,omitempty"`
	} `json:"hits"`
}

// MSearchQuery Multi Search query
type MSearchQuery struct {
	Header string // index name, document type
	Body   string // query related to the declared index
}

// MSearchResult Multi search result
type MSearchResult struct {
	Responses []SearchResult `json:"responses"`
}

var User1 = NewUser{
	ID:        int32(uuid.New().ID()),
	FirstName: "First thing first yo",
	LastName:  "Last",
	FullName:  "First Last",
	Username:  "UserOne",
	Bio:       "Hello darkness my old friend",
}

var User2 = NewUser{
	ID:        int32(uuid.New().ID()),
	FirstName: "Timburg",
	LastName:  "Last",
	FullName:  "Timburg Last",
	Username:  "UserTwo",
	Bio:       "I'm in love with a shooting star",
}

var User3 = NewUser{
	ID:        int32(uuid.New().ID()),
	FirstName: "Timbear",
	LastName:  "Last",
	FullName:  "Third Last",
	Username:  "UserThree",
	Bio:       "I'm sdkajfldsj sdifaef ok",
}

var User4 = NewUser{
	ID:        int32(uuid.New().ID()),
	FirstName: "Andrew",
	LastName:  "Last",
	FullName:  "Andrew Last",
	Username:  "Userfour",
	Bio:       "lorem ipsum ggg grag htfh",
}