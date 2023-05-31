package controllers

import (
	"Go_Project/models"
	"Go_Project/utils"
	"Go_Project/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateBookInput struct {
	Title      string `json:"title" binding:"required"`
	AuthorName string `json:"author_name" binding:"required"`
}

// @Summary get all items in the book list
// @ID get-all-books
// @Produce json
// @Success 200 {object} utils.Response
// @Router /get-all-books [get]
func GetAllBooks(c *gin.Context) {
	appG := utils.Gin{C: c}
	var books []models.Book
	models.DB.Find(books)
	appG.Response(http.StatusOK, e.SUCCESS, books)
}

// @Summary get a book item by ID
// @ID get-book-by-id
// @Produce json
// @Param id path string true "book_id"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /get-book-by-id/{id} [get]
func GetBooksById(c *gin.Context) {
	appG := utils.Gin{C: c}
	var book models.Book

	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {

		appG.Response(http.StatusBadRequest, e.SUCCESS, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, book)
}

// @Summary add a new item to the book list
// @ID create-book
// @Produce json
// @Param title body string true "Title"
// @Param author_name body string true "AuthorName"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /create-book [post]
func CreateBook(c *gin.Context) {
	appG := utils.Gin{C: c}
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {

		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	book := models.Book{Title: input.Title, AuthorName: input.AuthorName}
	models.DB.Create(&book)
	appG.Response(http.StatusOK, e.SUCCESS, book)
}
