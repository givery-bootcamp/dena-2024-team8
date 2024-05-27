package entities
import (
	"time"
)

type Post struct {
	Id int `json:"id"`
	UserId int `json:user_id`
	Title string `json:title`
	Body string `json:body`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:update_at`
	DeletedAt time.Time `json:deleted_at`
}
