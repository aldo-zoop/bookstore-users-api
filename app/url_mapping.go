package app

import (
	"github.com/aldo-zoop/bookstore-users-api/controller/ping"
	"github.com/aldo-zoop/bookstore-users-api/controller/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}
