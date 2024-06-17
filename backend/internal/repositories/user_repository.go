package repositories

import (
	"myapp/internal/entities"
	"time"
)

type User struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func ConvertUserRepositoryModelToEntity(user *User) *entities.User {
	return &entities.User{
		Id:        user.Id,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
