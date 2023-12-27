package models

type UserRights struct {
	Id         uint `json:"id" gorm:"primary_key"`
	UserId     uint `json:"user_id"`
	ElectionId uint `json:"election_id"`
	IsUsed     bool `json:"is_used"`

	User     User     `gorm:"foreignKey:UserId;references:Id"`
	Election Election `gorm:"foreignKey:ElectionId;references:Id"`
}

func (u *UserRights) TableName() string {
	return "user_rights"
}
