package middleware

import (
	"fmt"
	"net/http"
	"myapp/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// やること
// JWTValidをgin.HandlerFuncが返すようにする
// JWTValidを使って、JWTの検証を行う

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := config.JwtSecret
		if key == "" {
			// JWT_SECRETが設定されていない場合はエラーを返す
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET is not set"})
			ctx.Abort()
			return
		}
		tokenString, err := ctx.Cookie("jwt")
		if err != nil || tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		//tokenStringからpayloadを取り出す
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
		
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(key), nil
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			//claims["userId"]の値と型を確認
			fmt.Printf("%T, %v\n", claims["userId"], claims["userId"])
			floqtUserId, ok := claims["userId"].(float64)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": ""})
				ctx.Abort()
				return
			}
			ctx.Set("userId", int(floqtUserId))
			ctx.Next()
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			ctx.Abort()
			return
		}
	}
}