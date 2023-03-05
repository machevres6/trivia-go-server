package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/machevres6/trivia-go-server/models"
	"github.com/machevres6/trivia-go-server/database"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title": "Trivia Time Bitches",
		"Subtitle": "Facts for funtimes with friends!",
		"Facts": facts,
	})
}

// Create new fact view handler
func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title": "New Fact",
		"Subtitle": "Add a cool fact!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	// Create a new fact
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return NewFactView(c)
	}

	// Create a new fact in the database
	result := database.DB.Db.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}

	return ListFacts(c)
}

func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title": "Facts added successfully",
		"Subtitle": "Add more wonderful facts to the game",
	})
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.Status(fiber.StatusOK).Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact": fact,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
}

func EditFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	factEdit := database.DB.Db.Where("id = ?", id).First(&fact)
	if factEdit.Error != nil {
		return NotFound(c)
	}

	return c.Render("edit", fiber.Map{
		"Title": "Edit Fact",
		"Subtitle": "Edit your interesting fact",
		"Fact": fact,
	})
}

func UpdateFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	// Parsing the request body
	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	// Write updated values to the database
	factEdit := database.DB.Db.Where("id = ?", id).Updates(&fact)
	if factEdit.Error != nil {
		return EditFact(c)
	}

	return ShowFact(c)
}

func DeleteFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&fact)
	if result.Error != nil {
		return NotFound(c)
	}

	return ListFacts(c)
}