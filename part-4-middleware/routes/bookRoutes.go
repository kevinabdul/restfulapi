package routes

import(
	book "restfulalta/part-4-middleware/controllers/book"
	"restfulalta/part-4-middleware/middlewares"
)

func registerBookRoutes() {
	e.GET("/books", book.GetBooksController)

	e.GET("/books/:id", book.GetBookByIdController)

	e.POST("/books", book.AddBookController, middlewares.AuthenticateUser)

	e.PUT("/books/:id", book.EditBookController, middlewares.AuthenticateUser)

	e.DELETE("/books/:id", book.DeleteBookController, middlewares.AuthenticateUser)
}