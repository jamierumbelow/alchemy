package alchemy

import (
	"fmt"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/urfave/cli"
	"time"
)

// Run the main Alchemy tests
func Run(params Parameters, c *cli.Context) error {
	output := Output{}
	runKey := fmt.Sprintf("%d", time.Now().Unix())
	client := algoliasearch.NewClient(params.ConfigValues.AppID, params.ConfigValues.SecretKey)

	for _, indexName := range params.Indexes {
		tmpIndex, err := createTmpIndex(client, runKey, indexName)
		if err != nil {
			return err
		}
		defer tmpIndex.Delete()

		var fixtures []algoliasearch.Object
		err = ParseConfigFile(params.ConfigValues.FixturesPath, &fixtures)
		if err != nil {
			return err
		}

		_, err = addObjects(tmpIndex, fixtures)
		if err != nil {
			return err
		}

		var tests []TestCase
		err = ParseConfigFile(params.ConfigValues.TestsPath, &tests)
		for _, test := range tests {
			passes, err := test.Passes(tmpIndex)
			if !passes {
				output.Failure(test, err.Error())
			} else {
				output.Success(test)
			}
		}
	}

	return nil
}

func createTmpIndex(client algoliasearch.Client, runKey string, indexName string) (algoliasearch.Index, error) {
	originalIndex := client.InitIndex(indexName)
	tmpIndex := client.InitIndex(fmt.Sprintf("alchemy_%s_%s", runKey, indexName))

	originalSettings, err := originalIndex.GetSettings()
	if err != nil {
		return nil, err
	}

	_, err = tmpIndex.SetSettings(originalSettings.ToMap())
	if err != nil {
		return nil, err
	}

	return tmpIndex, nil
}

func addObjects(index algoliasearch.Index, fixtures []algoliasearch.Object) ([]algoliasearch.Object, error) {
	_, err := index.AddObjects(fixtures)
	return fixtures, err
}
