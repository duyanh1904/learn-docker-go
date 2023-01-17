package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	// select
	//queue := make(chan int)
	//done := make(chan bool)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		queue <- i
	//	}
	//	done <- true
	//}()
	//for {
	//	select {
	//	case v := <-queue:
	//		fmt.Println(v)
	//	case <-done:
	//		fmt.Println("done")
	//		return
	//	}
	//}

	//close channel

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
