package routes

import(
	book "restfulalta/part-3-code-structuring/controllers/book"
)

func registerBookRoutes() {
	e.GET("/books", book.GetBooksController)

	e.GET("/books/:id", book.GetBookByIdController)

	e.POST("/books", book.AddBookController)

	e.PUT("/books/:id", book.EditBookController)

	e.DELETE("/books/:id", book.DeleteBookController)
}