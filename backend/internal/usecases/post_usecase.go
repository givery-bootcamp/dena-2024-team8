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

func (u *PostUsecase) GetList(limit int, offset int) (
	[]*entities.Post, error) {
	return u.repository.List(nil, limit, offset)
}

func (u *PostUsecase) Get(id int) (*entities.Post, error) {
	return u.repository.Get(id)
}

func (u *PostUsecase) Delete(id int) error {
	return u.repository.Delete(id)
}
