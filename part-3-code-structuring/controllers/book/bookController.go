package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	book "restfulalta/part-3-code-structuring/services/book"
	models "restfulalta/part-3-code-structuring/models"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	fmt.Println("Inside getBooksController")
	books, err := book.GetBooks()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}

func GetBookByIdController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))

	targetBook, rowsAffected, err := book.GetBookById(targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong Book Id"})
	}
	return c.JSON(http.StatusOK, targetBook)
}

func AddBookController(c echo.Context) error {
	newBook := models.Book{}
	c.Bind(&newBook)

	res, err := book.AddBook(&newBook)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		Book models.Book
	}{Status: "succes", Message: "Book has been created!", Book: res})

}

func EditBookController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	newData := models.Book{}
	c.Bind(&newData)

	edittedBook, rowsAffected, err := book.EditBook(newData, targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong Book Id"})
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		Book models.Book
	}{Status: "succes", Message: "Book has been updated!", Book: edittedBook})
}

func DeleteBookController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	deleted, rowsAffected, err := book.DeleteBook(targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong Book Id"})
	}
	

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		Book models.Book
	}{Status: "succes", Message: "Book has been deleted!", Book: deleted})

}
