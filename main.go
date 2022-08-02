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

		image, saveImageErr := c.FormFile("image")

		if saveImageErr == nil {
			c.SaveFile(image, "./components/input.jpg")
		} else {
			return c.SendString("Error")
		}

		ComplierPython()

		file, readJsonErr := ioutil.ReadFile("output.json")
		if readJsonErr != nil {
			log.Fatal(readJsonErr)
		}

		var output []Outputjson

		jsonErr := json.Unmarshal(file, &output)

		if jsonErr != nil {
			fmt.Println(jsonErr)
		}

		removeInputErr := os.Remove("./components/input.jpg")
		if removeInputErr != nil {
			log.Fatal(removeInputErr)
		}

		removeOutputErr := os.Remove("output.json")
		if removeOutputErr != nil {
			log.Fatal(removeOutputErr)
		}

		return c.JSON(output)
	})

	log.Fatal(app.Listen(":8000"))
}
