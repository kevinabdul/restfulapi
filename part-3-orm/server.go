package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	gorm.Model
	Id     		int 	`json:"id" gorm:"primaryKey`
	Name   		string	`json:"name" form:"name"`
	Email 		string	`gorm:"unique" json:"email" form:"email"`
	Password 	string	`json:"password" form:"password"`
}

var (
	db *gorm.DB
)

func initDb() {
	err1 := godotenv.Load("../.env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var err2 error
	db, err2 = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err2 != nil {
		panic(err2)
	}

	db.AutoMigrate(&User{})
}

func main() {
	initDb()

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/users", getusers)

	e.GET("/users/:id", getUserById)

	e.POST("/users", addUser)

	e.PUT("/users/:id", editUser)

	e.DELETE("/users/:id", deleteUser)

	e.Start(":8000")
}

func getusers(c echo.Context) error {
	var users []User
	res := db.Find(&users)

	if res.Error != nil {
		return c.JSON(http.StatusBadRequest, 123)
	}
	return c.JSON(http.StatusOK, users)
}

func getUserById(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	var user User

	res := db.Find(&user, targetId)

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}
	return c.JSON(http.StatusOK, user)
}

func addUser(c echo.Context) error {
	newUser := User{}
	c.Bind(&newUser)

	db.Create(&newUser)
	
	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User User
	}{Status: "succes", Message: "User has been created!", User: newUser})

}

func editUser(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))

	newData := &User{}
	c.Bind(newData)
	
	targetUser := &User{}

	res := db.Where(`id = ?`, targetId).Find(&targetUser).Omit("password", "id").Updates(newData)
	
	if res.Error != nil {
		return c.JSON(http.StatusBadRequest, res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User User
	}{Status: "succes", Message: "User has been updated!", User: *targetUser})
}

func deleteUser(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	targetUser := &User{}
	res := db.Find(targetUser, targetId)

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}
	
	deleted := *targetUser
	db.Delete(targetUser, targetId)

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User User
	}{Status: "succes", Message: "User has been deleted!", User: deleted})

}
