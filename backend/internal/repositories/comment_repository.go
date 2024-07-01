package repositories

import (
	"log"
	"myapp/internal/entities"
	"time"

	"gorm.io/gorm"
)

type CommentRepository struct {
	Conn *gorm.DB
}

type Comment struct {
	Id        int       `json:"id"`
	UserId    int       `gorm:"foreignKey:UserId" json:"user_id"`
	PostId    int       `gorm:"foreignKey:PostId" json:"post_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"update_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewCommentRepository(conn *gorm.DB) *CommentRepository {
	return &CommentRepository{
		Conn: conn,
	}
}

func (r *CommentRepository) List(postId int, limit int, offset int) ([]*entities.Comment, error) {
	var comments []*Comment
	var result = r.Conn.Where("post_id = ?", postId).Limit(limit).Offset(offset).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	var entities []*entities.Comment
	for _, v := range comments {
		entities = append(entities, convertCommentRepositoryModelToEntity(v))
	}

	return entities, nil
}

func (r *CommentRepository) Create(postId int, body string, userId int) (*entities.Comment, error) {
	comment := Comment{
		PostId:    postId,
		Body:      body,
		UserId:    userId,
		DeletedAt: time.Date(9998, 12, 31, 23, 59, 59, 0, time.UTC),
	}
	// Check if the post exists
	var post Post
	var postResult = r.Conn.First(&post, postId)
	if postResult.Error != nil {
		log.Println("Post compatible with the postID is not found")
		return nil, postResult.Error
	}

	var result = r.Conn.Create(&comment)
	if result.Error != nil {
		return nil, result.Error
	}

	return convertCommentRepositoryModelToEntity(&comment), nil
}

func (r *CommentRepository) Update(commentId int, body string, userId int) (*entities.Comment, error) {
	// CreatedAt would not be updated
	comment := Comment{
		Id:        commentId,
		Body:      body,
		UserId:    userId,
		UpdatedAt: time.Now(),
	}
	// Check if the user is the owner of the comment
	var cc *Comment
	var commentResult = r.Conn.First(&cc, commentId)
	if commentResult.Error != nil {
		log.Println("Comment compatible with the commentID is not found")
		return nil, commentResult.Error
	}
	if comment.UserId != userId {
		log.Println("User is not the owner of the comment")
		return nil, gorm.ErrRecordNotFound
	}

	var result = r.Conn.Updates(&comment)
	if result.Error != nil {
		return nil, result.Error
	}

	return convertCommentRepositoryModelToEntity(&comment), nil
}

func (r *CommentRepository) Delete(commentId int, userId int) error {
	// Check if the user is the owner of the comment
	var comment Comment
	var commentResult = r.Conn.First(&comment, commentId)
	if commentResult.Error != nil {
		log.Println("Comment compatible with the commentID is not found")
		return commentResult.Error
	}
	if comment.UserId != userId {
		log.Println("User is not the owner of the comment")
		return gorm.ErrRecordNotFound
	}

	var result = r.Conn.Delete(&Comment{}, commentId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func convertCommentRepositoryModelToEntity(v *Comment) *entities.Comment {
	return &entities.Comment{
		Id:        v.Id,
		UserId:    v.UserId,
		PostId:    v.PostId,
		Body:      v.Body,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		DeletedAt: v.DeletedAt,
	}
}
