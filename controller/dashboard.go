package controller

import (
	"fmt"
	"go-election/database/connection"
	"go-election/models"

	"github.com/gofiber/fiber/v2"
)

func Dashboard(c *fiber.Ctx) error {
	db, sql, _ := connection.ConnectDB()
	defer sql.Close()

	var allElections []models.Election
	db.Select("id, name").
		Where("is_active = ?", true).
		Limit(3).
		Order("id").
		Find(&allElections)

	var election models.Election

	electionId := c.Query("election")
	if electionId != "" {
		db.Preload("Candidates").Where("id = ?", electionId).First(&election)
	} else {
		db.Preload("Candidates").Where("is_active = ?", true).Order("id DESC").First(&election)
	}

	var votes []models.Vote
	db.Preload("Candidate").Where("election_id = ?", election.Id).Find(&votes)

	if len(votes) == 0 {
		return c.Render("dashboard", fiber.Map{
			"AllElections": allElections,
			"Election":     election,
			"Message":      "No votes yet",
			"DataVote":     nil,
			"csrf":         c.Locals("token"),
		})
	}

	dataVote := map[string]uint{}

	for _, vote := range votes {
		dataVote[vote.Candidate.Leader] = dataVote[vote.Candidate.Leader] + 1
	}

	fmt.Println(election.Id, election.Name, votes[0].ElectionId)

	return c.Render("dashboard", fiber.Map{
		"AllElections": allElections,
		"Election":     election,
		"Message":      c.Locals("message"),
		"DataVote":     dataVote,
		"csrf":         c.Locals("token"),
	})
}
