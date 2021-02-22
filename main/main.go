package main

import (
	"fmt"
	"log"
	"strconv"

	data "automation-gmbh.com/datavis/Data"
	types "automation-gmbh.com/datavis/Types"
	"automation-gmbh.com/datavis/filehandling"
	"github.com/gofiber/fiber/v2"
)

// Anzahl an Helden
const (
	count = 20
)

func main() {
	//Standard aufrufe
	storage := data.NewStorage()

	fpath := "storage.json"

	err := storage.Load(fpath)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/api/file", func(c *fiber.Ctx) error {
		return c.JSON(filehandling.GiveFileAsStringArray())
	})

	app.Get("/api/heroes", func(c *fiber.Ctx) error {
		heros := storage.AllHeros()
		return c.JSON(heros)
	})

	app.Get("/api/hero/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return err
		}

		hero := storage.GetHero(int(id))

		if hero == nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.JSON(hero)
	})

	app.Put("/api/hero/:id", func(c *fiber.Ctx) error {

		idStr := c.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return err
		}

		hero := &types.Hero{}
		err = c.BodyParser(hero)
		if err != nil {
			return err
		}

		hero.Id = int(id)
		storage.UpdateHero(hero)

		err = storage.Save(fpath)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Delete("/api/hero/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return err
		}

		storage.RemoveHero(int(id))

		err = storage.Save(fpath)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/api/hero", func(c *fiber.Ctx) error {
		hero := &types.Hero{}
		err := c.BodyParser(hero)
		if err != nil {
			return err
		}

		id := storage.NextId()

		hero.Id = id

		storage.UpdateHero(hero)

		err = storage.Save(fpath)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	})

	log.Fatal(app.Listen(":3000"))

	message := data.Hello("Gladys")
	fmt.Print(message)

}
