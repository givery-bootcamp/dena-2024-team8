package entities

import (
	"time"
)

type Comment struct {
	Id        int       `json:"id"`
	UserId    int       `gorm:"foreignKey:UserId" json:"user_id"`
	PostId    int       `gorm:"foreignKey:PostId" json:"post_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
