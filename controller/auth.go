package controller

import (
	"go-election/database/connection"
	"go-election/models"
	"go-election/utils/jwt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Initialize the database connection

func Login(c *fiber.Ctx) error {
	var db, conn, _ = connection.ConnectDB()
	defer conn.Close()

	// cek remembered token from cookie
	lastToken, _ := jwt.Get(c)
	if lastToken != "" {
		return c.Redirect("/?message=Already logged in", fiber.StatusFound)
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
	if err != nil || !comparePasswords(request.Password, user.Password) {
		return c.Redirect("/?message=invalid credentials", fiber.StatusFound)
	}

	token, _ := jwt.GenerateToken(user.Id, user.Name, user.Role)

	jwt.Save(c, token, user.Id)

	if user.Role == "admin" {
		return c.Redirect("/dashboard?message=login success", fiber.StatusFound)
	}

	return c.Redirect("/?message=login success", fiber.StatusFound)
}

func Logout(c *fiber.Ctx) error {
	// Set header Cache-Control
	c.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")

	// Set header Pragma
	c.Set("Pragma", "no-cache")

	// Set header Expires
	expirationTime := time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC)
	c.Set("Expires", expirationTime.Format(time.RFC1123))

	jwt.Delete(c)
	c.Method("GET")
	return c.Redirect("/?message=logout success", fiber.StatusFound)
}

// Helper function to compare hashed passwords
func comparePasswords(password string, hashedPassword string) bool {
	// Implement your password comparison logic here
	// Example: return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
	return password == hashedPassword
}
