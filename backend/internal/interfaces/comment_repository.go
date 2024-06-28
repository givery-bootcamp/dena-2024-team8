package interfaces

import (
	"myapp/internal/entities"
)

type CommentRepository interface {
	Create(postId int, body string, userId int) (*entities.Comment, error)
	Update(commentId int, body string, userId int) (*entities.Comment, error)
	Delete(commentId int, userId int) error
}
