package models

type Vote struct {
	Id           uint `json:"id" gorm:"primary_key"`
	UserRightsId uint `json:"user_rights_id"`
	CandidateId  uint `json:"candidate_id"`

	UserRights UserRights `gorm:"foreignKey:UserRightsId;references:Id"`
	Candidate  Candidate  `gorm:"foreignKey:CandidateId;references:Id"`
}

func (v *Vote) TableName() string {
	return "votes"
}
