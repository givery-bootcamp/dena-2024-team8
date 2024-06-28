package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"myapp/internal/entities"
	"myapp/internal/external"
	"myapp/internal/interfaces"
	"strconv"
	"strings"
)

type PostUsecase struct {
	repository interfaces.PostRepository
}

func NewPostUsecase(r interfaces.PostRepository) *PostUsecase {
	return &PostUsecase{
		repository: r,
	}
}

func (u *PostUsecase) GetList(limit int, offset int) (
	[]*entities.Post, error) {
	return u.repository.List(nil, limit, offset)
}

func (u *PostUsecase) Get(id int) (*entities.Post, error) {
	return u.repository.Get(id)
}

func (u *PostUsecase) Create(title, body string, userId int) (*entities.Post, error) {
	return u.repository.Create(title, body, userId)
}

func (u *PostUsecase) Update(title, body string, userId, postId int) (*entities.Post, error) {
	return u.repository.Update(title, body, userId, postId)
}

func (u *PostUsecase) Search(query string) ([]int, error) {
	var searchQuery string
	if query == "" {
		// Query is empty, search for all documents
		searchQuery = `{
			"query": {
				"match_all": {}
			}
		}`
	} else {
		searchQuery = fmt.Sprintf(`{
			"query": {
				"multi_match": {
					"query": "%s",
					"fields": ["title", "body"]
				}
			}
		}`, query)
	}

	var buf strings.Builder
	buf.WriteString(searchQuery)

	res, err := external.ES.Search(
		external.ES.Search.WithContext(context.Background()),
		external.ES.Search.WithIndex("posts"),
		external.ES.Search.WithBody(strings.NewReader(buf.String())),
		external.ES.Search.WithTrackTotalHits(true),
		external.ES.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Search error: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %v", err)
			return nil, err
		} else {
			log.Printf("Error: %s: %s", res.Status(), e["error"].(map[string]interface{})["reason"])
			return nil, errors.New("error")
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %v", err)
		return nil, err
	}

	postsIds := []int{}

	hits, ok := r["hits"].(map[string]interface{})
	if !ok || hits["hits"] == nil {
		// No hits found, handle empty result case
		return postsIds, nil
	}

	for _, hit := range hits["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		fmt.Printf("source: %v\n", source["id"])
		id, _ := source["id"].(string) // Type assertion with fallback
		intId, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Error parsing id: %v", err)
			continue
		}
		postsIds = append(postsIds, intId)
	}

	return postsIds, nil
}
