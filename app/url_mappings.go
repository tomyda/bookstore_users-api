package app

import (
	"github.com/tomyda/bookstore_users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
