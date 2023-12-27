package jwt

import (
	"go-election/models"
	"testing"
)

var (
	tokenStr   string
	err        error
	userClaims *models.UserClaims
)

func TestGenerateToken(t *testing.T) {
	tokenStr, err = GenerateToken(1, "test")
	if err != nil {
		t.Errorf("Failed to generate token: %v\n", err)
		return
	}
}

func TestVerifyToken(t *testing.T) {
	userClaims, err = VerifyToken(tokenStr)
	if err != nil {
		t.Errorf("Failed to verify token: %v\n", err)
		return
	}
}
