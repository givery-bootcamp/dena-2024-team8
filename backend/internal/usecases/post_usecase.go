package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type PostUsecase struct {
	repository interfaces.PostRepository
}

func NewPostUsecase(r interfaces.PostRepository) *PostUsecase {
	return &PostUsecase{
		repository: r,
	}
}

func (u *PostUsecase) GetList() (
	[]*entities.Post, error) {
	return u.repository.List()
}

func (u *PostUsecase) Get(id int) (*entities.Post, error) {
	return u.repository.Get(id)
}
