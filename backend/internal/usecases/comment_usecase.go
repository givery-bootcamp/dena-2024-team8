package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type CommentUsecase struct {
	repository interfaces.CommentRepository
}

func NewCommentUsecase(r interfaces.CommentRepository) *CommentUsecase {
	return &CommentUsecase{
		repository: r,
	}
}

func (u *CommentUsecase) List(postId int, limit int, offset int) ([]*entities.Comment, error) {
	return u.repository.List(postId, limit, offset)
}

func (u *CommentUsecase) Create(postId int, body string, userId int) (*entities.Comment, error) {
	return u.repository.Create(postId, body, userId)
}

func (u *CommentUsecase) Update(commentId int, body string, userId int) (*entities.Comment, error) {
	return u.repository.Update(commentId, body, userId)
}

func (u *CommentUsecase) Delete(commentId int, userId int) error {
	return u.repository.Delete(commentId, userId)
}
