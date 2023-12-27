package jwt

import (
	"fmt"
	"go-election/config"
	"go-election/models"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey []byte

func init() {
	data, _ := config.LoadConfig()
	secretKey = []byte(data.JwtSecret)
}

// GenerateToken membuat token JWT dengan klaim khusus
func GenerateToken(userId int, username string) (string, error) {
	// Set waktu kadaluarsa token (misalnya, 1 jam dari sekarang)
	expirationTime := time.Now().Add(time.Hour * 1).Unix()

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
