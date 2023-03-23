package routers

import (
	"challenge06/controller"
	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine) *gin.Engine {
	router.POST("/books", controller.CreateBook)       // Buat buku
	router.GET("/books", controller.GetAllBooks)       // Ambil semua buku
	router.GET("/books/:id", controller.GetBook)       // Ambil buku berdasarkan id
	router.PUT("/books/:id", controller.UpdateBook)    // Update buku
	router.DELETE("/books/:id", controller.DeleteBook) // Hapus buku

	return router
}
