package main

import (
	"fmt"
	"github.com/blevesearch/bleve/v2"
)

const IndexRootPath = "indexes"

func CreateIndex(indexName string) error {
	indexPath := fmt.Sprintf("%s.%s", IndexRootPath, indexName)
	mapping := bleve.NewIndexMapping()
	_, err := bleve.New(indexPath, mapping)
	if err != nil {
		return err
	}
	return nil
}

func InsertData(indexName string, dataId string, dataOwner string, data string) error {
	indexPath := fmt.Sprintf("%s.%s", IndexRootPath, indexName)
	dataStruct := struct {
		Id    string
		Owner string
		Data  string
	}{Id: dataId, Owner: dataOwner, Data: data}

	index, err := bleve.Open(indexPath)
	if err != nil {
		return err
	}
	err = index.Index(dataId, dataStruct)
	if err != nil {
		return err
	}

	err = index.Close()
	return err
}

func SearchData(indexName string, searchQuery string) (result string, err error) {
	indexPath := fmt.Sprintf("%s.%s", IndexRootPath, indexName)
	index, err := bleve.Open(indexPath)
	if err != nil {
		return "", err
	}
	query := bleve.NewQueryStringQuery(searchQuery)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		return "", err
	}
	// TODO: add limit for result
	err = index.Close()
	if err != nil {
		return "", err
	}
	return searchResult.String(), nil
}

func DeleteData(indexName string, dataId string) error {
	indexPath := fmt.Sprintf("%s.%s", IndexRootPath, indexName)
	index, err := bleve.Open(indexPath)
	if err != nil {
		return err
	}

	err = index.Delete(dataId)
	if err != nil {
		return err
	}
	err = index.Close()
	return err
}
