package search

import (
	"github.com/meilisearch/meilisearch-go"
)

type MeiliSearch struct {
	Client *meilisearch.Client
}

func (ms *MeiliSearch) Setup(host, masterKey string) {
	ms.Client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   host,
		APIKey: masterKey,
	})
}

func (ms *MeiliSearch) AddDocument(index string, document any) error {
	_, err := ms.Client.Index(index).UpdateDocuments(document)
	return err
}

func (ms *MeiliSearch) SearchDocument(index, text string) (interface{}, error) {
	res, err := ms.Client.Index(index).Search(text, &meilisearch.SearchRequest{
		MatchingStrategy: "all",
	})

	if err != nil {
		return nil, err
	}

	return res.Hits, nil
}
