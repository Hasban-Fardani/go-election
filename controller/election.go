package controller

import (
	"go-election/database/connection"
	"go-election/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Election(c *fiber.Ctx) error {
	db, conn, _ := connection.ConnectDB()
	defer conn.Close()

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

func AdminElections(c *fiber.Ctx) error {
	return c.Render("admin_elections", nil)
}

type Data struct {
	Draw            int               `json:"draw"`
	RecordsTotal    int64             `json:"recordsTotal"`
	RecordsFiltered int64             `json:"recordsFiltered"`
	Data            []models.Election `json:"data"`
}

func APIElections(c *fiber.Ctx) error {
	db, conn, _ := connection.ConnectDB()
	defer conn.Close()

	// Membaca parameter yang dikirimkan oleh datatable
	draw, _ := strconv.Atoi(c.Query("draw"))
	start, _ := strconv.Atoi(c.Query("start"))
	length, _ := strconv.Atoi(c.Query("length"))
	search := c.Query("search[value]")
	log.Println(draw, start, length, search)

	// Membuat variabel untuk menyimpan data
	var data Data
	var elections []models.Election
	var count int64
	var total int64

	// Menghitung jumlah total data di tabel elections
	db.Model(&models.Election{}).Count(&total)

	// Menyaring data berdasarkan kata kunci pencarian jika ada
	if search != "" {
		db.Limit(length).
			Order("id asc").
			Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%").
			Find(&elections).
			Count(&count)
	} else {
		db.
			Offset(start).
			Order("id asc").
			Find(&elections).
			Count(&count)
	}

	// Mengisi data dengan nilai yang sesuai
	data.Draw = draw
	data.RecordsTotal = total
	data.RecordsFiltered = count
	data.Data = elections

	// log.Println(count, total, search, elections, data.Data)
	// Mengirimkan response JSON dengan data yang dibutuhkan oleh datatable
	return c.JSON(data)
}
