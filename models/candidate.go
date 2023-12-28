package models

type Candidate struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Leader      string `json:"leader"`
	Vice        string `json:"vice"`
	Image       string `json:"image"`
	Description string `json:"description"`
	ElectionId  uint   `json:"election_id"`

	Election Election `gorm:"foreignKey:ElectionId;references:Id"`
}

func (c *Candidate) TableName() string {
	return "candidates"
}
