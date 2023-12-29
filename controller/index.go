package controller

import (
	"go-election/database/connection"
	"go-election/models"
	"go-election/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

type Data []map[string]any

func Index(c *fiber.Ctx) error {
	message := c.Locals("message")

	csrfToken := c.Locals("token").(string)
	token, _ := jwt.Get(c)

	if token == "" {
		return c.Render("index", fiber.Map{
			"csrf":       csrfToken,
			"LoggedIn":   false,
			"UserRights": []models.UserRights{},
			"Message":    message,
		})
	}

	db, sqlDb, _ := connection.ConnectDB()
	defer sqlDb.Close()

	userId := jwt.GetUserId(c)

	// get user rights
	var userRights []models.UserRights
	db.Where("user_id = ?", userId).
		Preload("Election", "is_active = 1").
		Order("election_id").
		Find(&userRights)

	return c.Render("index", fiber.Map{
		"csrf":       csrfToken,
		"LoggedIn":   true,
		"UserRights": userRights,
		"Message":    message,
	})
}
