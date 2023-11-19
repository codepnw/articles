package handlers

import (
	"github.com/codepnw/articles/database"
	"github.com/codepnw/articles/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title":    "Hello Golang",
		"Subtitle": "coding with me.",
		"Facts":    facts,
	})
}

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title": "New Fact",
		"Subtitle": "Add a fact!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return ConfirmationView(c)
}

func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title": "Fact added successfully",
		"Subtitle": "Add more wonderful facts to the list!",
	})
}