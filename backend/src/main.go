package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Outputjson struct {
	Class       string  `json:"class"`
	Probability float64 `json:"probability"`
}

func Complier() Outputjson {

	cmd := exec.Command("python", "app.py", "--model", "0", "--img", "input.jpg")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}

	file, readJsonErr := ioutil.ReadFile("output.json")
	if readJsonErr != nil {
		log.Fatal(readJsonErr)
	}

	var output []Outputjson

	jsonErr := json.Unmarshal(file, &output)

	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	removeInputErr := os.Remove("input.jpg")
	if removeInputErr != nil {
		log.Fatal(removeInputErr)
	}

	removeOutputErr := os.Remove("output.json")
	if removeOutputErr != nil {
		log.Fatal(removeOutputErr)
	}

	return output[0]
}

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Get("/test", func(c *fiber.Ctx) error {

		type Data struct {
			Name string
		}

		data := Data{
			Name: "Test",
		}

		return c.JSON(data)
	})

	app.Post("/post-image", func(c *fiber.Ctx) error {

		image, saveImageErr := c.FormFile("image")

		if saveImageErr == nil {
			c.SaveFile(image, "input.jpg")
		} else {
			return c.SendString("Error")
		}

		return c.JSON(Complier())
	})

	app.Post("post-url-image", func(c *fiber.Ctx) error {

		imageURL, err := url.Parse(c.FormValue("url"))
		if err != nil {
			panic(err)
		}

		res, err := http.Get(imageURL.String())
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		file, err := os.Create("input.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = io.Copy(file, res.Body)
		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(Complier())
	})

	log.Fatal(app.Listen(":4000"))
}
