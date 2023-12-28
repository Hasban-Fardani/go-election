package controller

import (
	"go-election/database/connection"
	"go-election/models"
	"go-election/utils/jwt"
	"go-election/utils/response"
	"go-election/utils/validator"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Login(c *fiber.Ctx) (err error) {
	store := session.New()
	sess, _ := store.Get(c)

	db, sqlDb, _ := connection.ConnectDB()
	defer sqlDb.Close()

	var request models.UserRequest
	if err = c.BodyParser(&request); err != nil {
		log.Println("login | err: ", err.Error())
		sess.Set("message", "Invalid request")
		sess.Save()
		return response.JsonError(c, fiber.StatusBadRequest, err)
	}

	if validator.EmptyStr(request.Email, request.Password) {
		log.Println("login | err: Invalid request ", request.Email, request.Password)
		sess.Set("message", "Invalid request")
		sess.Save()
		return c.RedirectBack("/")
	}

	var user models.User
	err = db.Where("email = ?", request.Email).First(&user).Error

	if err != nil {
		log.Println("login | err: ", err.Error())
		sess.Set("message", "Invalid email")
		sess.Save()
		return c.RedirectBack("/")
	}

	if request.Password != user.Password {
		log.Println("login | err: ", err.Error())
		sess.Set("message", "Invalid password")
		sess.Save()
		return c.RedirectBack("/")
	}

	token, _ := jwt.GenerateToken(user.Id, user.Name)

	jwt.Save(c, token)

	return c.RedirectBack("/")
}

func Logout(c *fiber.Ctx) error {
	jwt.Delete(c)
	return c.RedirectBack("/")
}

func APILogin(c *fiber.Ctx) (err error) {
	db, sqlDb, _ := connection.ConnectDB()
	defer sqlDb.Close() // close connection at the end

	request := new(models.UserRequest)
	if err = c.BodyParser(&request); err != nil {
		return response.JsonError(c, fiber.StatusBadRequest, err)
	}

	if validator.EmptyStr(request.Email, request.Password) {
		return c.RedirectBack("/")
	}

	var user models.User
	err = db.Where("email = ?", request.Email).First(&user).Error

	if err != nil {
		return response.JsonError(c, fiber.StatusInternalServerError, err)
	}

	if user.Password != request.Password {
		return response.JsonError(c, fiber.StatusUnauthorized, err)
	}

	token, _ := jwt.GenerateToken(user.Id, user.Email)

	jwt.Save(c, token)

	return c.JSON(fiber.Map{"token": token, "message": "login success", "user": user})
}

func APICheckLogin(c *fiber.Ctx) error {
	token := jwt.Get(c)
	if token == "" {
		return c.JSON(fiber.Map{"message": "unauthorized"})
	}
	return c.JSON(fiber.Map{"token": token, "message": "authorized"})
}

func APILogout(c *fiber.Ctx) error {
	jwt.Delete(c)
	return c.JSON(fiber.Map{"message": "logout success"})
}
