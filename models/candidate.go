package models

type Candidate struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ElectionId  uint   `json:"election_id"`

	Election Election `gorm:"foreignKey:ElectionId;references:Id"`
}

func (c *Candidate) TableName() string {
	return "candidates"
}
