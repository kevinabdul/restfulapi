package controllers

import (
	"net/http"
	"strconv"

	user "restfulalta/part-4-middleware/services/user"
	models "restfulalta/part-4-middleware/models"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, err := user.GetUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserByIdController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	targetUser, rowsAffected, err := user.GetUserById(targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}
	return c.JSON(http.StatusOK, targetUser)
}

func AddUserController(c echo.Context) error {
	newUser := models.User{}
	c.Bind(&newUser)

	res, err := user.AddUser(&newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User models.UserAPI
	}{Status: "success", Message: "User has been created!", User: res})

}

func EditUserController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))

	newData:= models.User{}
	c.Bind(&newData)

	edittedUser, rowsAffected, err := user.EditUser(newData, targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User models.UserAPI
	}{Status: "success", Message: "User has been updated!", User: edittedUser})
}

func DeleteUserController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	rowsAffected, err := user.DeleteUser(targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}
	

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
	}{Status: "success", Message: "User has been deleted!"})

}

func LoginUserController(c echo.Context) error {
	loggingUser := &models.User{}
	c.Bind(loggingUser)

	token, err := user.LoginUser(loggingUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		Token string
	}{Status: "success", Message: "You are logged in!", Token: token})
}
