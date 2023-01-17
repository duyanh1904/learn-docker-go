package controllers

import (
	"fmt"
	"github.com/duyanh1904/learn-docker-go/mariab_db"
	model "github.com/duyanh1904/learn-docker-go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type UserController struct{}

var hmacSampleSecret []byte

func (u UserController) Retrieve(c *gin.Context) {
	db := mariab_db.InitDb()
	var user []model.User
	users := model.GetUsers(db, &user)

	c.JSON(http.StatusCreated, gin.H{"message": "Get data users", "Data": map[string]interface{}{"data": users}})
}

func (u UserController) GenToken(c *gin.Context) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2022, 11, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString, err)

	c.String(http.StatusOK, "Token: ", tokenString)
}

//
//func (u UserController) Create(c *gin.Context) {
//	db := mariab_db.InitDb()
//	var user []model.User
//	users := model.CreateUser(db, &user)
//
//	c.JSON(http.StatusCreated, gin.H{"message": "Create success", "Data": map[string]interface{}{"data": user}})
//	//c.JSON(http.StatusCreated, gin.H{"message": "Create success", "Data": map[string]interface{}{"data": users}})
//}

func (u UserController) Update(c *gin.Context) {
	var user []model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := c.PostForm("name")
	age := c.PostForm("age")
	c.JSON(http.StatusOK, gin.H{"name": name, "age": age, "user": user})
}
