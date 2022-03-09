package main

import (
	"api-with-mongodb/controllers/user_controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	user_controller.BuildRoutes(router)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
