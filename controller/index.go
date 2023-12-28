package controller

import (
	"go-election/database/connection"
	"go-election/models"
	"go-election/utils/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Index(c *fiber.Ctx) error {
	store := session.New()
	sess, _ := store.Get(c)
	message := sess.Get("message")

	csrfToken := c.Locals("token").(string)
	token := jwt.Get(c)
	if token == "" {
		return c.Render("index", fiber.Map{
			"csrf":       csrfToken,
			"LoggedIn":   false,
			"Candidates": []models.Candidate{},
			"Message":    message,
		})
	}

	db, sqlDb, _ := connection.ConnectDB()
	defer sqlDb.Close()

	var candidates []models.Candidate
	db.Preload("Election").Find(&candidates)

	return c.Render("index", fiber.Map{
		"csrf":       csrfToken,
		"LoggedIn":   true,
		"Candidates": candidates,
		"Message":    message,
	})
}
