package models

type Election struct {
	Id            uint   `json:"id" gorm:"primary_key"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	NumberOfVotes uint   `json:"number_of_votes"`
	IsActive      bool   `json:"is_active"`

	Candidates []Candidate `gorm:"foreignKey:ElectionId;references:Id"`
}
