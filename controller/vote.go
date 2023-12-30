package controller

import (
	"go-election/database/connection"
	"go-election/models"
	"go-election/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func UserVote(c *fiber.Ctx) (err error) {
	// init database
	db, sqlDb, _ := connection.ConnectDB()
	defer sqlDb.Close()

	// find candidate from param
	candidateId := c.Params("id")
	var candidate models.Candidate
	err = db.Where("id = ?", candidateId).
		Preload("Election").
		Find(&candidate).
		Error

	// find election from candidate
	var election models.Election
	err = db.Debug().
		Preload("Candidates").
		Where("id = ?", candidate.ElectionId).
		First(&election).
		Error
	if err != nil {
		return err
	}

	// get user from jwt
	user := jwt.GetUser(c, db)

	// get user rights
	// where user_id == user.Id &&
	//       election_id == candidate.ElectionId &&
	//       is_used == false
	var userRights models.UserRights
	err = db.Debug().
		Where("user_id = ?", user.Id).
		Where("election_id = ?", election.Id).
		Where("is_used = ?", false).
		Preload("User").
		Find(&userRights).
		Error
	if err != nil || userRights.Id == 0 {
		return c.Redirect("/?message=you don't have rights", fiber.StatusFound)
	}

	// check if user has voted
	var voteCount int64
	err = db.Debug().
		Where("user_rights_id = ?", userRights.Id).
		Count(&voteCount).
		Error

	// if user has voted
	if voteCount > 0 {
		c.Locals("message", "user has voted")
		return c.Redirect("/?message=user has voted", fiber.StatusFound)
	}

	// if user has not voted
	err = db.Create(&models.Vote{
		CandidateId:  candidate.Id,
		UserRightsId: userRights.Id,
		ElectionId:   election.Id,
	}).Error

	db.Model(&userRights).Update("is_used", true)

	return c.Redirect("/", fiber.StatusFound)
}
