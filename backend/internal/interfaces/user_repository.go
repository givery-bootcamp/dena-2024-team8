package interfaces

import (
	"myapp/internal/entities"
)

type UserRepository interface {
	VerifyUser(username, password string) (*entities.User, error)
	Get(id int) (*entities.User, error)
}
