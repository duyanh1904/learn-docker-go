package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	model "github.com/duyanh1904/learn-docker-go/models"
	database "github.com/duyanh1904/learn-docker-go/mongo_db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm/utils"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Bird struct {
	Species     string
	Description string
}

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

func (h HealthController) JsonArrays(c *gin.Context) {
	birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	err := json.Unmarshal([]byte(birdJson), &birds)
	if err != nil {
		return
	}
	fmt.Printf("Birds : %+v", birds)

	c.JSON(http.StatusOK, birds)
}

func (h HealthController) MakeChannel(c *gin.Context) {
	queue := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}
		close(queue)
	}()

	for value := range queue {
		fmt.Println(value)
	}

	c.String(http.StatusOK, "Working!")
}

type csvUploadInput struct {
	CsvFile *multipart.FileHeader `form:"file" binding:"required"`
}

func (h HealthController) ImportCsv(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	if err := c.ShouldBind(&file); err != nil {
		log.Fatal(err)
	}

	log.Println(file.Filename)

	err = c.SaveUploadedFile(file, "kafka-logs/"+utils.ToString(rand.Int()))
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

type BookRecord struct {
	Title        string
	Author       string
	Price        int32
	Discount     int32
	Organization string
}

func (h HealthController) Insert(c *gin.Context) {
	// Use a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(i int) {
			// Decrement the wait group count when the goroutine finishes
			defer wg.Done()

			// Insert the record here. This could be a database call, for example.
			insert(i)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	c.JSON(http.StatusAccepted, "success")
}

func insert(i int) {
	mongo := database.ConnectDB()
	collection := database.GetCollection(mongo, "test_perform")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := strconv.Itoa(i)

	birdPayload := model.Post{
		ID:      primitive.NewObjectID(),
		Title:   "post.Title" + s,
		Article: "post.Article",
	}
	result, err := collection.InsertOne(ctx, birdPayload)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(birdPayload)
	fmt.Println(result)
}
