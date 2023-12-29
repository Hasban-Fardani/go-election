package controller

import (
	"go-election/database/connection"
	"go-election/models"

	"github.com/gofiber/fiber/v2"
)

func Election(c *fiber.Ctx) error {
	db, sqlDb, _ := connection.ConnectDB()
	defer sqlDb.Close()

	// overide method before redirect
	c.Method("GET")

	// get id param
	electionId := c.Params("id")

	// get candidates from electionId
	var candidates []models.Candidate
	db.Debug().
		Where("election_id = ?", electionId).
		Find(&candidates)

	csrfToken := c.Locals("token").(string)
	return c.Render("election", fiber.Map{
		"Candidates": candidates,
		"csrf":       csrfToken,
	})
}
