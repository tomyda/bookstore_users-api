package app

import (
	"github.com/tomyda/bookstore_users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:user-id", controllers.GetUser)
}
