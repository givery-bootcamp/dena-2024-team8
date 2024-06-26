package usecases

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/interfaces"
	"os"

	"github.com/golang-jwt/jwt"
)

type UserUsecase struct {
	repository interfaces.UserRepository
}

func NewUserUsecase(r interfaces.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: r,
	}
}

func generateJWT(userId int) (string, error) {
	key := os.Getenv("JWT_SECRET") // 環境変数から取得する場合
	if key == "" {
		return "", errors.New("JWT_SECRET is not set")
	}
	// JWTトークンのクレーム（ペイロード）を設定
	claims := jwt.MapClaims{
		"userId": userId,
	}

	// 新しいJWTトークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// シークレットキーで署名し、トークン文字列を生成
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *UserUsecase) VerifyUserAndGenerateJWT(username, password string) (*entities.User, *string, error) {
	var user, err = u.repository.VerifyUser(username, password)
	if err != nil {
		return nil, nil, err
	}
	if user != nil {
		userId := user.Id
		jwtToken, err := generateJWT(userId)
		if err != nil {
			return nil, nil, err
		}
		return user, &jwtToken, nil
	} else {
		return user, nil, nil
	}
}

// func (u *PostUsecase) Get(id int) (*entities.Post, error) {
// 	return u.repository.Get(id)
// }
	func (u *UserUsecase) Get(id int) (*entities.User, error) {
		return u.repository.Get(id)
	}
