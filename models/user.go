package models

type User struct {
	Id              uint   `json:"id" gorm:"primary_key"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	RememberedToken string `json:"-"`
	Password        string `json:"-"`

	UserRights []UserRights `gorm:"foreignKey:UserId;references:Id"`
}

func (u *User) TableName() string {
	return "users"
}
