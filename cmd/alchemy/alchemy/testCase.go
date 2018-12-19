package alchemy

import (
	"fmt"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

// TestCase represents a single test against a single index
type TestCase struct {
	query           interface{} `json:"query"`
	ExpectedResults []string    `json:"expectedResults"`
}

// Name returns a human-readable name for the test
func (test TestCase) Name() string {
	return test.Query()
}

// Query parses and returns the query string
func (test TestCase) Query() string {
	query, ok := test.query.(string)
	if ok {
		return query
	}

	return query
}

// Passes returns true when the test passes on a given index, false (with error message) otherwise
func (test TestCase) Passes(index algoliasearch.Index) (bool, error) {
	results, err := index.Search(test.Query(), nil)
	fmt.Print(results)
	return false, err
}
