package app

import (
	ping_controller "github.com/tomyda/bookstore_users-api/controllers/ping"
	"github.com/tomyda/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping_controller.Ping)

	router.GET("/users/:user-id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
