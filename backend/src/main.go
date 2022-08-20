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

var access_tok = ""

func ErrFunc(err error, c *gin.Context) {
	data := Data{
		Text: err.Error(),
	}
	c.JSON(0, data)
}

func Complier(c *gin.Context) {

	cmd := exec.Command("python3", "app.py", "--model", "0", "--img", "input.jpg")
	err := cmd.Run()

	if err != nil {
		ErrFunc(err, c)
		return
	}

	file, readJsonErr := ioutil.ReadFile("output.json")
	if readJsonErr != nil {
		ErrFunc(readJsonErr, c)
		return
	}

	var output []Outputjson

	jsonErr := json.Unmarshal(file, &output)

	if jsonErr != nil {
		ErrFunc(jsonErr, c)
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
		ErrFunc(err, c)
		return
	}
}

func uploadurl(c *gin.Context) {
	imgurl := c.PostForm("url")
	res, err := http.Get(imgurl)
	if err != nil {
		ErrFunc(err, c)
		return
	}
	defer res.Body.Close()

	file, err := os.Create("input.jpg")
	if err != nil {
		ErrFunc(err, c)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		ErrFunc(err, c)
		return
	}
	Complier(c)
}

func getjwt(c *gin.Context) {

	res := c.Request.Header.Get("Access")

	if res != api_key {
		return
	} else {
		token, err := CreateJWT()
		if err != nil {
			return
		}
		access_tok = token
	}
}

func getfooddata(c *gin.Context) {
	name := c.Param("food")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://username:password@cluster0.h1omugq.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		ErrFunc(err, c)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			ErrFunc(err, c)
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
		ErrFunc(err, c)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		ErrFunc(err, c)
	}
	c.String(http.StatusOK, "%s\n", jsonData)
}

func getjwtwithkey(c *gin.Context) {

	res := c.Param("key")

	if res != api_key {
		return
	} else {
		token, err := CreateJWT()
		if err != nil {
			return
		}
		access_tok = token
	}
}

func Validate(next func(c *gin.Context)) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if ValidateJWT(access_tok) {
			access_tok = ""
			next(c)

		} else {
			c.JSON(http.StatusUnauthorized, "can't access")
		}
	})
}

func main() {
	app := gin.Default()

	app.GET("/test", Validate(servercheck))
	app.POST("post-image", Validate(uploadimage))
	app.POST("/post-url", Validate(uploadurl))
	app.GET("/data/:food", Validate(getfooddata))
	app.GET("/access-key", getjwt)
	app.GET("/access-key/:key", getjwtwithkey)

	app.Run(":5000")
}
