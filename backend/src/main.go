package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Text string `json:"text"`
}

type Outputjson struct {
	Class       string  `json:"class"`
	Probability float64 `json:"probability"`
}

func Complier(c *gin.Context) {

	cmd := exec.Command("python3", "app.py", "--model", "0", "--img", "input.jpg")
	err := cmd.Run()

	if err != nil {
		data := Data{
			Text: err.Error(),
		}
		c.JSON(0, data)
		return
	}

	file, readJsonErr := ioutil.ReadFile("output.json")
	if readJsonErr != nil {
		data := Data{
			Text: readJsonErr.Error(),
		}
		c.JSON(0, data)
		return
	}

	var output []Outputjson

	jsonErr := json.Unmarshal(file, &output)

	if jsonErr != nil {
		data := Data{
			Text: jsonErr.Error(),
		}
		c.JSON(0, data)
		return
	}

	c.JSON(http.StatusOK, output[0])

}

func servercheck(c *gin.Context) {

	data := Data{
		Text: "Golang",
	}
	c.JSON(http.StatusOK, data)
}

func uploadimage(c *gin.Context) {
	c.Header("Content-Type", "image/jpeg")
	file, err := c.FormFile("image")
	if err == nil {
		c.SaveUploadedFile(file, "input.jpg")
		Complier(c)
	} else {
		data := Data{
			Text: err.Error(),
		}
		c.JSON(0, data)
		return
	}
}

func uploadurl(c *gin.Context) {
	imgurl := c.PostForm("url")
	res, err := http.Get(imgurl)
	if err != nil {
		data := Data{
			Text: err.Error(),
		}
		c.JSON(0, data)
		return
	}
	defer res.Body.Close()

	file, err := os.Create("input.jpg")
	if err != nil {
		data := Data{
			Text: err.Error(),
		}
		c.JSON(0, data)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		data := Data{
			Text: err.Error(),
		}
		c.JSON(0, data)
		return
	}
	Complier(c)
}

func main() {
	app := gin.Default()

	app.GET("/test", servercheck)
	app.POST("post-image", uploadimage)
	app.POST("/post-url", uploadurl)

	app.Run(":5000")
}
