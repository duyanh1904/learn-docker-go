package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	model "github.com/duyanh1904/learn-docker-go/models"
	database "github.com/duyanh1904/learn-docker-go/mongoDb"
	kafkaRun "github.com/duyanh1904/learn-docker-go/worker"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"time"
)

type PostController struct{}

func (p PostController) CreatePost(c *gin.Context) {
	var DB = database.ConnectDB()

	bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
	post := new(model.Post)

	inputJSON := string(bodyAsByteArray)
	err := json.Unmarshal([]byte(inputJSON), &post)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	var postCollection = database.GetCollection(DB, "Posts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postPayload := model.Post{
		ID:      primitive.NewObjectID(),
		Title:   post.Title,
		Article: post.Article,
	}

	result, err := postCollection.InsertOne(ctx, postPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}

func (p PostController) RunKafka(c *gin.Context) {
	kafkaRun.Run()
	c.JSON(http.StatusCreated, gin.H{"message": "Push message success"})
}
