package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"myapp/internal/entities"
	"time"
)

type PostRepository struct {
	Conn *gorm.DB
}

type Post struct {
	Id        int       `json:"id"`
	UserId    int       `json:user_id`
	Title     string    `json:title`
	Body      string    `json:body`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:update_at`
	DeletedAt time.Time `json:deleted_at`
}

func NewPostRepository(conn *gorm.DB) *PostRepository {
	return &PostRepository{
		Conn: conn,
	}
}

func (r *PostRepository) List() ([]*entities.Post, error) {
	var obj []Post
	result := r.Conn.Order("id desc").Find(&obj)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", obj)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	pes := convertSlices(obj, convertPostRepositoryModelToEntity)
	return pes, nil
}

// convertSlices は []T を []U へ変換します
func convertSlices[T, U any](srcList []T, convertFunc func(*T) U) []U {
	var result []U
	for _, v := range srcList {
		result = append(result, convertFunc(&v))
	}
	return result
}

func convertPostRepositoryModelToEntity(v *Post) *entities.Post {
	return &entities.Post{
		Id:        v.Id,
		UserId:    v.UserId,
		Title:     v.Title,
		Body:      v.Body,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		DeletedAt: v.DeletedAt,
	}
}
