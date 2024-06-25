package repositories

import (
	"errors"
	"myapp/internal/entities"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

type User struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (r *UserRepository) VerifyUser(username, password string) (*entities.User, error) {
	user := &User{}
	result := r.Conn.Where("name = ? AND password = ?", username, password).First(user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return ConvertUserRepositoryModelToEntity(user), nil
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
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

// この関数は、UserRepository型のメソッドとして定義されています。FindByIPassという名前の関数です。
// この関数は、与えられたユーザー名とパスワードを使用して、データベース内のユーザーを検索します。まず、与えられたパスワードをハッシュ化し、そのハッシュ値を出力します。次に、usersテーブルから指定されたユーザー名に一致するユーザーを検索します。
// もし検索結果がエラーでない場合、ユーザーのパスワードと与えられたパスワードを比較します。もしパスワードが一致しない場合、エラーを返します。一致する場合は、ConvertUserRepositoryModelToEntity関数を使用して、データベースモデルをエンティティオブジェクトに変換して返します。
// この関数は、ユーザーが存在しない場合やエラーが発生した場合には、適切な値（nilやエラーオブジェクト）を返します。
func (r *UserRepository) FindByIPass(username string, password string) (*entities.User, error) {
	// SignUpの時にここの処理を使って、パスワードをハッシュ化してDBに保存する
	// hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	// println(string(hash))

	var user User
	result := r.Conn.Where("name = ?", username, password).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}
	// ユーザのパスワードと与えられたパスワードを比較
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return ConvertUserRepositoryModelToEntity(&user), nil
}