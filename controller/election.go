package controller

import (
	"go-election/database/connection"
	"go-election/models"
	"go-election/utils/response"

	"github.com/gofiber/fiber/v2"
)

func GetElections(c *fiber.Ctx) error {
	db, sqlDb, _ := connection.ConnectDB()
	defer sqlDb.Close()

	var elections []models.Election
	smt := db.Preload("Candidates")

	// filtering
	if c.Query("active") == "true" {
		smt = smt.Where("is_active = ?", true)
	}

	if c.Query("active") == "false" {
		smt = smt.Where("is_active = ?", false)
	}

	if c.Query("name") != "" {
		smt = smt.Where("name LIKE ?", "%"+c.Query("name")+"%")
	}

	if c.Query("description") != "" {
		smt = smt.Where("description LIKE ?", "%"+c.Query("description")+"%")
	}

	smt.Find(&elections)
	return c.JSON(fiber.Map{"status": "success", "data": elections})
}

func GetElectionById(c *fiber.Ctx) error {
	db, sqlDb, _ := connection.ConnectDB()
	defer sqlDb.Close()

	var election models.Election

	if err := db.Where("id = ?", c.Params("id")).First(&election).Error; err != nil {
		return response.JsonError(c, fiber.StatusNotFound, err)
	}

	return c.JSON(fiber.Map{"status": "success", "data": election})
}

// TODO: create election for admin
func CreateElection(c *fiber.Ctx) error {
	return nil
}
