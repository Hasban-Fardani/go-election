package controller

import (
	"go-election/database/connection"
	"go-election/models"
	"go-election/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	db, sqlDb, _ := connection.ConnectDB()
	defer sqlDb.Close()

	// cek remembered token from cookie
	lastToken, _ := jwt.Get(c)
	if lastToken != "" {
		c.Redirect("/?message=Already logged in", fiber.StatusFound)
	}

	// overide method before redirect
	c.Method("GET")

	var request models.UserRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Redirect("/?message=invalid request", fiber.StatusFound)
	}

	var user models.User
	err := db.Where("email = ?", request.Email).
		First(&user).
		Error

	// cek password and error
	if err != nil || request.Password != user.Password {
		return c.Redirect("/?message=invalid credentials", fiber.StatusFound)
	}

	token, _ := jwt.GenerateToken(user.Id, user.Name, user.Role)

	jwt.Save(c, token, user.Id)

	user.RememberedToken = token
	db.Save(&user)

	if user.Role == "admin" {
		return c.Redirect("/dashboard?message=login success", fiber.StatusFound)
	}

	return c.Redirect("/?message=login success", fiber.StatusFound)
}

func Logout(c *fiber.Ctx) error {
	jwt.Delete(c)
	c.Method("GET")
	return c.Redirect("/?message=logout success", fiber.StatusFound)
}
