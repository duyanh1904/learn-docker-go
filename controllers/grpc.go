package controllers

import (
	pb "github.com/duyanh1904/learn-docker-go/pb/grpc-server/v1"
	"github.com/gin-gonic/gin"
)

type GrpcController struct{}
type Person struct {
}

func (h GrpcController) GetPerson(c *gin.Context) {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	c.JSON(200, p)
}
