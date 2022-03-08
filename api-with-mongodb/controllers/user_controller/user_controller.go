package user_controller

import (
	"api-with-mongodb/models"
	"api-with-mongodb/services/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	endpoint = "/users"
)

func Routers(router *gin.Engine) {
	router.GET(endpoint, GetAll)
	router.GET(endpoint+"/:id", GetById)
	router.POST(endpoint, InsertOne)
	router.PUT(endpoint+"/:id", UpdateOne)
	router.DELETE(endpoint+"/:id", DeleteOne)
}

func GetAll(c *gin.Context) {

	users, err := user_service.Read()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": "No users found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetById(c *gin.Context) {

	id := c.Param("id")

	user, err := user_service.ReadByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func InsertOne(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user_service.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "User created successfully"})
}

func UpdateOne(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := user_service.Update(user, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User updated successfully"})
}

func DeleteOne(c *gin.Context) {

	id := c.Param("id")
	if err := user_service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User deleted successfully"})
}
