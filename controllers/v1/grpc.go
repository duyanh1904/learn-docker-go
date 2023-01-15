package controllers

import (
	"github.com/gin-gonic/gin"
	pb "github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
	"net/http"
)

type grpcController struct{}
type Person struct {
}

func (h grpcController) Status(c *gin.Context) {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	c.String(http.StatusOK, "Working!")
}
