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

func BuildRoutes(router *gin.Engine) {
	router.GET(endpoint, FindMany)
	router.GET(endpoint+"/:id", FindOneByID)
	router.POST(endpoint, InsertOne)
	router.PUT(endpoint+"/:id", UpdateOne)
	router.DELETE(endpoint+"/:id", DeleteOne)
}

func FindMany(c *gin.Context) {

	users, err := user_service.FindMany()
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

func FindOneByID(c *gin.Context) {

	id := c.Param("id")

	user, err := user_service.FindOneByID(id)
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

	if err := user_service.InsertOne(user); err != nil {
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
	if err := user_service.UpdateOne(user, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User updated successfully"})
}

func DeleteOne(c *gin.Context) {

	id := c.Param("id")
	if err := user_service.DeleteOne(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User deleted successfully"})
}
