package models

import "github.com/golang-jwt/jwt"

type User struct {
	Id              int    `json:"id" gorm:"primary_key"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"-"`
	RememberedToken string `json:"-"`

	UserRights []UserRights `gorm:"foreignKey:UserId;references:Id"`
}

func (u *User) TableName() string {
	return "users"
}

// user claims for jwt
type UserClaims struct {
	UserId   int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// user request for login
type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	CRSF     string `json:"csrf_token" form:"csrf_token"`
}
