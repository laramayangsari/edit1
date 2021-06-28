package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id    int
	Age   int
	Email string
	Name  string
}

func main() {
	e := echo.New()

	///routing
	e.GET("/user/:id/:age", UserController)
	e.POST("/user", RegisterController)
	e.Start(":8000")
}

//controller
func UserController(e echo.Context) error {

	id, _ := strconv.Atoi(e.Param("id"))

	age, _ := strconv.Atoi(e.Param("age"))

	search := e.QueryParam("search")

	short := e.QueryParam("short")

	user := User{Id: id, Age: age, Email: "academy@alterra.id", Name: "Alta"}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"user":   user,
		"search": search,
		"short":  short,
	})
}

func RegisterController(c echo.Context) error {
	email := c.FormValue("email")
	name := c.FormValue("name")
	age := c.FormValue("age")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"email": email,
		"name":  name,
		"age":   age,
	})
}
