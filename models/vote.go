package models

import (
	"errors"

	"gorm.io/gorm"
)

type Vote struct {
	Id           uint `json:"id" gorm:"primary_key"`
	UserRightsId uint `json:"user_rights_id"`
	CandidateId  uint `json:"candidate_id"`
	ElectionId   uint `json:"election_id"`

	UserRights UserRights `gorm:"foreignKey:UserRightsId;references:Id"`
	Candidate  Candidate  `gorm:"foreignKey:CandidateId;references:Id"`
}

func (v *Vote) TableName() string {
	return "votes"
}

func (v *Vote) BeforeCreate(tx *gorm.DB) (err error) {
	// check if the user rights is not used
	if v.UserRights.IsUsed {
		return errors.New("User is already voted")
	}
	return
}
