package elasticsearch

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Match struct {
	Match map[string]interface{} `json:"match,omitempty"`
}

type MyQuery struct {
	Query struct {
		Bool struct {
			Must   []*Match `json:"must"`
			Should []*Match `json:"should"`
		} `json:"bool"`
	} `json:"query"`
}

func TestBuildElasticSearchQuery(t *testing.T) {
	q := buildQuery()
	byteArr, _ := json.Marshal(q)
	fmt.Println(string(byteArr))
	// example output:
	// {
	//   query: {
	//    bool: {
	// 			must:[
	// 				{ match: { "MyField1": 123 }},
	// 				{ match: { "MyField2": "some-string" }},
	// 			should:[
	// 				{ match: { "MyField3": time.Now() }},
	// 				{ match: { "MyField4": true }},
	// 			}
	//   }
	// }
}

func buildQuery() *MyQuery {
	q := MyQuery{}
	q.Query.Bool.Must = []*Match{
		buildMatch("MyField1", 123),
		buildMatch("MyField2", "some-string"),
	}
	q.Query.Bool.Should = []*Match{
		buildMatch("MyField3", time.Now()),
		buildMatch("MyField4", true),
	}
	return &q
}

func buildMatch(name string, value interface{}) *Match {
	return &Match{
		Match: map[string]interface{}{name: value},
	}
}
