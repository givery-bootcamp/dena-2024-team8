package interfaces

import (
	"myapp/internal/entities"
)

type PostRepository interface {
	Get(id int) (*entities.Post, error)
	List(id *int, limit int, offset int) ([]*entities.Post, error)
	Delete(id int) error
}
