package middleware

import (
	"ecommerce/dto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(c *fiber.Ctx) error {
	c.Locals("auth", false)
	// Authorization başlığını al
	authHeader := c.Get("Authorization")
	// Tokeni al
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	if tokenString != "" {
		// Tokeni kontrol et
		token, err := jwt.Parse(
			tokenString, func(token *jwt.Token) (interface{}, error) {
				// JWT tokenin doğrulanacağı gizli anahtar veya RSA PublicKey gibi bir yapıyı döndürün.
				// Burada tokenin doğrulanacağı algoritmayı ve gerekli anahtarı sağlamanız gerekmektedir.
				// Örneğin:
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Geçersiz token"})
		}

		// Token doğrulandı, kullanıcı bilgisine erişebilirsiniz
		claims := token.Claims.(jwt.MapClaims)

		//userID := claims["user_id"].(string)
		group := claims["group"].(float64)
		email := claims["email"].(string)
		id := claims["id"].(float64)

		user := dto.User{Group: group, Email: email, Id: id}
		c.Locals("user", user)

	} else {
		user := dto.User{Group: 0, Email: "", Id: 0}
		c.Locals("user", user)
	}

	return c.Next()
}
