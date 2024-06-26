package repositories

import (
	"errors"
	"fmt"
	"myapp/internal/entities"
	"time"

	"gorm.io/gorm"
)

type PostRepository struct {
	Conn *gorm.DB
}

type Post struct {
	Id        int       `json:"id"`
	UserId    int       `json:user_id`
	User      User      `gorm:"foreignKey:UserId" json:user`
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

func (r *PostRepository) Get(id int) (*entities.Post, error) {
	posts, err := r.List(&id, 1, 0)
	if err != nil {
		return nil, err
	}
	if len(posts) == 0 {
		return nil, errors.New("not found")
	}
	return posts[0], nil
}

func (r *PostRepository) Create(title, body string, userId int) (*entities.Post, error) {
	post := Post{
		Title:     title,
		Body:      body,
		UserId:    userId,
		DeletedAt: time.Date(9998, 12, 31, 23, 59, 59, 0, time.UTC),
	}
	var result = r.Conn.Create(&post)

	if result.Error != nil {
		return nil, result.Error
	}
	return convertPostRepositoryModelToEntity(&post, &post.User), nil
}

func (r *PostRepository) List(id *int, limit int, offset int) ([]*entities.Post, error) {
	var obj []Post
	var result *gorm.DB
	if id != nil {
		result = r.Conn.Preload("User").Where("id = ?", id).Order("id desc").Find(&obj)
	} else {
		result = r.Conn.Preload("User").Order("id desc").Limit(limit).Offset(offset).Find(&obj)
	}
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", obj)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []*entities.Post{}, nil
		}
		return nil, result.Error
	}
	pes := convertSlices(obj, func(v *Post) *entities.Post {
		return convertPostRepositoryModelToEntity(v, &v.User)
	})

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

func convertPostRepositoryModelToEntity(v *Post, u *User) *entities.Post {
	return &entities.Post{
		Id:        v.Id,
		UserId:    v.UserId,
		User:      *ConvertUserRepositoryModelToEntity(u),
		Title:     v.Title,
		Body:      v.Body,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		DeletedAt: v.DeletedAt,
	}
}
