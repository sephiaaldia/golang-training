package main

import (
	"net/http"
	// "strconv"
	"github.com/labstack/echo"
)

type User struct{
	Id int `json: "id" form: "id"`
	Name string `json: "name" form: "name"`
	Email string `json: "email" form: "email"`
	Password string `json: "password" form: "password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages" : "success get all users",
		"users" : users,
	})
}

func GetUserController(c echo.Context) error {
	
}

func DeleteUserController(c echo.Context) error {
	
}

func UpdateUserController(c echo.Context) error {
	
}

func CreateUserController(c echo.Context) error {
	user := User {}
	c.Bind(&user)
	
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages" : "success create user",
		"user" : user,
	})
}

func challenge(){
	e := echo.New()
	e.GET("/users", GetUserController)
	e.GET("/users", CreateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}