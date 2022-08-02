package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

type Outputjson struct {
	Class       string  `json:"class"`
	Probability float64 `json"probability"`
}

func ComplierPython() {

	cmd := exec.Command("python", "app.py", "--model", "0", "--img", "./components/input.jpg")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}

	fmt.Printf("Complied")
}

func main() {

	app := fiber.New()

	app.Post("/postimage", func(c *fiber.Ctx) error {

		image, err := c.FormFile("image")

		if err == nil {
			c.SaveFile(image, "./components/input.jpg")
		} else {
			return c.SendString("Error")
		}

		return c.SendString("Image is uploaded")
	})

	app.Patch("/complie", func(c *fiber.Ctx) error {
		ComplierPython()
		return c.SendString("Finished")
	})

	app.Get("/result", func(c *fiber.Ctx) error {

		file, e := ioutil.ReadFile("output.json")
		if e != nil {
			log.Fatal(e)
		}

		var output []Outputjson

		err := json.Unmarshal(file, &output)

		if err != nil {
			fmt.Println(err)
		}

		return c.JSON(output)
	})

	app.Delete("/reset", func(c *fiber.Ctx) error {

		err := os.Remove("./components/input.jpg")
		if err != nil {
			log.Fatal(err)
		}

		e := os.Remove("output.json")
		if err != nil {
			log.Fatal(e)
		}

		return c.SendString("Deleted")
	})

	log.Fatal(app.Listen(":8000"))
}
