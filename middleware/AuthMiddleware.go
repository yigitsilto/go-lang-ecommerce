package middleware

import (
	model "ecommerce/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	c.Set("auth", false)
	// Authorization başlığını al
	authHeader := c.GetHeader("Authorization")
	// Tokeni al
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	if tokenString != "" {
		// Tokeni kontrol et
		token, err := jwt.Parse(
			tokenString, func(token *jwt.Token) (interface{}, error) {
				// JWT tokenin doğrulanacağı gizli anahtar veya RSA PublicKey gibi bir yapıyı döndürün.
				// Burada tokenin doğrulanacağı algoritmayı ve gerekli anahtarı sağlamanız gerekmektedir.
				// Örneğin:
				return []byte("4ruYgZPR6DuhlMZU0kTUBkNMzuE5e1KFh0jNvMdUyG8lo9xq10g7iChetp9escvx"), nil
			},
		)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz token"})
			return
		}

		// Token doğrulandı, kullanıcı bilgisine erişebilirsiniz
		claims := token.Claims.(jwt.MapClaims)

		//userID := claims["user_id"].(string)
		group := claims["group"].(float64)
		email := claims["email"].(string)

		user := model.User{Group: group, Email: email}
		c.Set("user", user)

	} else {
		user := model.User{Group: 0, Email: ""}
		c.Set("User", user)
	}

	c.Next()
}
