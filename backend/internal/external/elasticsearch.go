package external

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"myapp/internal/entities"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

var ES *elasticsearch.Client

func InitElasticSearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200",
		},
	}
	var err error
	ES, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	syncDataToElasticSearch()
}

func syncDataToElasticSearch() {
	var posts []entities.Post
	if err := DB.Find(&posts).Error; err != nil {
		log.Fatalf("failed to get posts: %v", err)
	}

	for _, post := range posts {
		// JSONエスケープするためにencoding/jsonパッケージを使用
		postJson, err := json.Marshal(map[string]string{
			"title": post.Title,
			"body":  post.Body,
		})
		if err != nil {
			log.Fatalf("Error marshaling post ID=%d: %s", post.Id, err)
		}

		req := esapi.IndexRequest{
			Index:      "posts",
			DocumentID: fmt.Sprintf("%d", post.Id),
			Body:       strings.NewReader(string(postJson)),
			Refresh:    "true",
		}

		res, err := req.Do(context.Background(), ES)
		if err != nil {
			log.Fatalf("Error indexing document ID=%d: %s", post.Id, err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("Error indexing document ID=%d: %s", post.Id, res.String())
		} else {
			log.Printf("Successfully indexed document ID=%d", post.Id)
		}
	}
}
