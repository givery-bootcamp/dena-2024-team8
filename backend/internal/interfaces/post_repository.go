package interfaces

import (
	"myapp/internal/entities"
)

type PostRepository interface {
	Get(id int) (*entities.Post, error)
	List(id *int, limit int, offset int) ([]*entities.Post, error)
	Create(title, body string, userId int) (*entities.Post, error)
	Update(title, body string, userId, postId int) (*entities.Post, error)
	Delete(id int) error
}
