package main

import (
	"fmt"
	"log"

	"example.com/data"
	"example.com/filehandling"
	"example.com/types"
	"github.com/gofiber/fiber/v2"
)

// Anzahl an Helden
const (
	count = 20
)

func main() {
	//Standard aufrufe
	data.Main()
	filehandling.Main()

	app := fiber.New()
	// app Foo

	app.Get("/foo", func(c *fiber.Ctx) error {
		fmt.Print("Go Sendet jetzt einen String an /foo")
		return c.SendString("Von go gesendeter String")
	})

	//app api/Heroes
	a := make([]types.Hero, count) // Erstellt ein Array mit Helden
	for i := 0; i < count; i++ {
		a[i] = data.GiveHero()
	}

	app.Get("/api/heroes", func(c *fiber.Ctx) error {
		return c.JSON(a)
	})

	app.Get("/api/file", func(c *fiber.Ctx) error {
		return c.JSON(filehandling.GiveFileAsStringArray())
	})

	log.Fatal(app.Listen(":3000"))

	message := data.Hello("Gladys")
	fmt.Print(message)

}
