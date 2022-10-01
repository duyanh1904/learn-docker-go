package controllers

import (
	"github.com/duyanh1904/learn-docker-go/mariabDB"
	model "github.com/duyanh1904/learn-docker-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (u UserController) Retrieve(c *gin.Context) {
	db := mariabDB.InitDb()
	var user []model.User
	users := model.GetUsers(db, &user)

	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": users}})
}
