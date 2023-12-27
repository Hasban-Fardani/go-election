package models

type Election struct {
	Id            uint   `json:"id" gorm:"primary_key"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	NumberOfVotes uint   `json:"number_of_votes"`
	IsActive      bool   `json:"is_active"`

	Candidates []Candidate  `gorm:"foreignKey:ElectionId;references:Id"`
	UserRights []UserRights `gorm:"foreignKey:ElectionId;references:Id"`
}

func (e *Election) TableName() string {
	return "elections"
}
