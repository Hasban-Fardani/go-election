package jwt

import (
	"fmt"
	"go-election/config"
	"go-election/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var secretKey []byte

func init() {
	data, _ := config.LoadConfig()
	secretKey = []byte(data.JwtSecret)
}

// GenerateToken membuat token JWT dengan klaim khusus
func GenerateToken(userId int, username string) (string, error) {
	// Set waktu kadaluarsa token
	expirationTime := time.Now().Add(time.Duration(config.Data.JwtExpireMinute) * time.Minute).Unix()

	// Membuat token dengan klaim khusus
	claims := &models.UserClaims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	// Membuat token JWT dengan HS256 algoritma
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token dengan secret key
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

// VerifyToken memverifikasi token JWT dan mengembalikan klaim jika valid
func VerifyToken(tokenString string) (*models.UserClaims, error) {
	// Parsing token dengan secret key
	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Memeriksa apakah parsing berhasil dan token valid
	if err != nil {
		return nil, err
	}

	// Memeriksa tipe klaim dan mengembalikan klaim khusus
	if claims, ok := token.Claims.(*models.UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func Save(c *fiber.Ctx, token string) {
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Minute * time.Duration(config.Data.JwtExpireMinute))
	cookie.HTTPOnly = true

	c.Cookie(cookie)
}

func Delete(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
}

func Get(c *fiber.Ctx) (token string) {
	// token = c.Get("Authorization")
	token = c.Cookies("token")
	if _, err := VerifyToken(token); err != nil {
		Delete(c)
		token = ""
	}
	return
}
