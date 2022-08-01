package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/post-image", func(c *fiber.Ctx) error {
		image, err := c.FormFile("image")
		if err == nil {
			c.SaveFile(image, fmt.Sprintf("./components/%s", image.Filename))
		} else {
			return c.SendString("Error")
		}
		return c.SendString("Image has uploaded.")
	})

	app.Post("/post-image-url", func(c *fiber.Ctx) error {

		imageURL, err := url.Parse(c.FormValue("url"))
		if err != nil {
			panic(err)
		}

		res, err := http.Get(imageURL.String())
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		file, err := os.Create(fmt.Sprintf("./components/%s.jpg", imageURL.Hostname()))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = io.Copy(file, res.Body)
		if err != nil {
			log.Fatal(err)
		}

		return c.SendString("Image has uploaded.")
	})

	log.Fatal(app.Listen(":8080"))
}
