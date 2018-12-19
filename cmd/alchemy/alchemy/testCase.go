package alchemy

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/jamierumbelow/alchemy/pkg/utils"
)

// TestCase represents a single test against a single index
type TestCase struct {
	Query           QueryInfo `json:"query"`
	ExpectedResults []string  `json:"expectedResults"`
}

// QueryInfo represents a test query
type QueryInfo struct {
	Query   string `json:"query"`
	Filters string `json:"filters"`
}

func (query QueryInfo) String() string {
	queryStr := fmt.Sprintf("\"%s\"", query.Query)

	if len(query.Filters) > 0 {
		queryStr = queryStr + fmt.Sprintf(" [%s]", query.Filters)
	}

	return queryStr
}

func (query QueryInfo) MarshalJSON() ([]byte, error) {
	if len(query.Filters) > 0 {
		return json.Marshal(query)
	}

	return json.Marshal(query.Query)
}

func (query *QueryInfo) UnmarshalJSON(data []byte) error {
	switch data[0] {
	case '"':
		var queryStr string
		if err := json.Unmarshal(data, &queryStr); err != nil {
			return err
		}
		query.Query = queryStr

	case '{':
		content := struct {
			Query   string `json:"query"`
			Filters string
		}{}

		if err := json.Unmarshal(data, &content); err != nil {
			return err
		}

		query.Query = content.Query
		query.Filters = content.Filters

	default:
		return errors.New("Unsupported query type")
	}

	return nil
}

// Name returns a human-readable name for the test
func (test TestCase) Name() string {
	return test.Query.String()
}

// Passes returns true when the test passes on a given index, false (with error message) otherwise
func (test TestCase) Passes(index algoliasearch.Index) (bool, error) {
	var params algoliasearch.Map

	if len(test.Query.Filters) > 0 {
		params = algoliasearch.Map{
			"filters": test.Query.Filters,
		}
	}

	results, err := index.Search(test.Query.Query, params)
	if err != nil {
		return false, err
	}

	var objectIDs []string
	for _, hit := range results.Hits {
		if str, ok := hit["objectID"].(string); ok {
			objectIDs = append(objectIDs, str)
		}
	}

	if utils.Equal(objectIDs, test.ExpectedResults) {
		return true, nil
	}

	return false, errors.New("Expected results weren't returned")
}
