package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func getfooddata(c *gin.Context) {
	name := c.Param("food")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://username:password@cluster0.h1omugq.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("what-the-food").Collection("menu")
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: name}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", name)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	c.String(http.StatusOK, "%s\n", jsonData)
}

func main() {
	app := gin.Default()

	app.GET("/test", servercheck)
	app.POST("post-image", uploadimage)
	app.POST("/post-url", uploadurl)
	app.GET("/data/:food", getfooddata)

	app.Run(":5000")
}
