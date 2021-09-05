package routes

import(
	user "restfulalta/part-3-code-structuring/controllers/user"
)

func registerUserRoutes() {
	e.GET("/users", user.GetUsersController)

	e.POST("/users", user.AddUserController)

	e.GET("/users/:id", user.GetUserByIdController)

	e.PUT("/users/:id", user.EditUserController)

	e.DELETE("/users/:id", user.DeleteUserController)
	
}

