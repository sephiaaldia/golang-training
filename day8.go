package main

import (
	// "fmt"
	"net/http"
	"github.com/labstack/echo"
	// "strconv"
)

// type User struct{
// 	Id int
// 	Name string
// 	Email string
// }

func UserSearchController(c echo.Context) error {
	match := c.QueryParam("match")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"match" : match,
		"result" : []string{"adi", "aan", "asif"},
	})
}

// func GetUserController(c echo.Context)error{
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	user := User{Id: id,Name:"Ismail", Email:"ismail@alterra.id"}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"user" : user, 
// 	})
// }

// func GetUser(c echo.Context)error{
// 	user := User{Name:"Ismail", Email:"ismail@alterra.id"}
// 	return c.JSON(http.StatusOK, user)

// }

func day8()  {
	e := echo.New()
	// e.GET("/", HelloController)
	//  e.GET("/user/:id", GetUserController)
	e.GET("/user", UserSearchController)
	e.Start(":8000")

}

// func HelloController(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello World!")
// }