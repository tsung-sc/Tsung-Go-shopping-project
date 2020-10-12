package models

import (
	"fmt"

	"github.com/olivere/elastic/v7"
)

var EsClient *elastic.Client

func init() {
	EsClient, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Println(err)
	}
}
