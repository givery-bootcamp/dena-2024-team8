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
			"http://127.0.0.1:9200",
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
			"id":    fmt.Sprintf("%d", post.Id),
			"title": post.Title,
			"body":  post.Body,
		})
		if err != nil {
			log.Fatalf("Error marshaling post ID=%d: %s", post.Id, err)
		}

		sendCreateRequest("posts", fmt.Sprintf("%d", post.Id), "true", postJson)
	}
}

func SendCreatePostRequest(post entities.Post) {
	postJson, err := json.Marshal(map[string]string{
		"id":    fmt.Sprintf("%d", post.Id),
		"title": post.Title,
		"body":  post.Body,
	})
	if err != nil {
		log.Fatalf("Error marshaling post ID=%d: %s", post.Id, err)
	}

	sendCreateRequest("posts", fmt.Sprintf("%d", post.Id), "true", postJson)
}

func SendUpdatePostRequest(post entities.Post) {
	postJson, err := json.Marshal(map[string]string{
		"id":    fmt.Sprintf("%d", post.Id),
		"title": post.Title,
		"body":  post.Body,
	})
	if err != nil {
		log.Fatalf("Error marshaling post ID=%d: %s", post.Id, err)
	}

	sendUpdateRequest("posts", fmt.Sprintf("%d", post.Id), "true", postJson)
}

func SendDeletePostRequest(id int) {
	sendDeleteRequest("posts", fmt.Sprintf("%d", id), "true")
}

func sendCreateRequest(index, documentID, refresh string, body []byte) {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       strings.NewReader(string(body)),
		Refresh:    refresh,
	}

	res, err := req.Do(context.Background(), ES)
	if err != nil {
		log.Fatalf("Error indexing document ID=%s: %s", documentID, err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error indexing document ID=%s: %s", documentID, res.String())
	} else {
		log.Printf("Successfully indexed document ID=%s", documentID)
	}
}

func sendUpdateRequest(index, documentID, refresh string, body []byte) {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       strings.NewReader(string(body)),
		Refresh:    refresh,
	}

	res, err := req.Do(context.Background(), ES)
	if err != nil {
		log.Fatalf("Error updating document ID=%s: %s", documentID, err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error updating document ID=%s: %s", documentID, res.String())
	} else {
		log.Printf("Successfully updated document ID=%s", documentID)
	}
}

func sendDeleteRequest(index, documentID, refresh string) {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: documentID,
		Refresh:    refresh,
	}

	res, err := req.Do(context.Background(), ES)
	if err != nil {
		log.Fatalf("Error deleting document ID=%s: %s", documentID, err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error deleting document ID=%s: %s", documentID, res.String())
	} else {
		log.Printf("Successfully deleted document ID=%s", documentID)
	}
}
