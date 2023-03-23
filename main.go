package main

import (
	"challenge06/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setting port
	PORT := ":8000"

	// Jalankan server
	initRouters := gin.Default()
	initRouters = routers.BookRouter(initRouters)

	err := initRouters.Run(PORT)
	if err != nil {
		panic(err)
	}
}
