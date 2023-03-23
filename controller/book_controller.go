package controller

import (
	"challenge06/entity"
	"challenge06/repository"
	"challenge06/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateBook(ctx *gin.Context) {
	// Isi variable newBook dari request
	var newBook entity.Book
	errBind := ctx.ShouldBindJSON(&newBook)
	if errBind != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid json",
		})
		return
	}

	// Panggil service untuk buat buku
	errService := service.CreateBook(newBook)
	if errService != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errService.Error(),
		})
		return
	}

	ctx.String(http.StatusCreated, "Created")
}

func GetAllBooks(ctx *gin.Context) {
	// Panggil service ambil semua buku
	err := service.GetAllBooks()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, repository.DatabaseBook)
}

func GetBook(ctx *gin.Context) {
	// Ambil id dari url
	id := ctx.Param("id")
	cvtId, errCvt := strconv.Atoi(id)
	if errCvt != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Input bukan tipe nomor",
		})
		return
	}

	// Panggil service untuk ambil buku berdasarkan id
	book, errService := service.GetBook(cvtId)
	if errService != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errService,
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func UpdateBook(ctx *gin.Context) {
	// Ambil id dari url
	id := ctx.Param("id")
	cvtId, errCvt := strconv.Atoi(id)
	if errCvt != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Input bukan tipe nomor",
		})
		return
	}

	// Ambil data json dari request
	var newBook entity.Book
	errJson := ctx.ShouldBindJSON(&newBook)
	if errJson != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid json",
		})
		return
	}

	// Panggil service update buku dengan id yang ingin diganti dan data buku baru
	errService := service.UpdateBook(cvtId, newBook)
	if errService != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errService.Error(),
		})
		return
	}

	ctx.String(http.StatusOK, "Updated")
}

func DeleteBook(ctx *gin.Context) {
	// Ambil id dari url
	id := ctx.Param("id")
	cvtId, errCvt := strconv.Atoi(id)
	if errCvt != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Input bukan tipe nomor",
		})
		return
	}

	// Panggil service untuk hapus data berdasarkan id
	errService := service.DeleteBook(cvtId)
	if errService != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errService.Error(),
		})
		return
	}

	ctx.String(http.StatusOK, "Deleted")
}
