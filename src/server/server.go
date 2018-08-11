package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ResultResponse : Create, Read 결과 반환용
type ResultResponse struct {
	Message string
}

func index(c echo.Context) error {
	// 일단 이것만 문자열로 전달
	result := "You are reached to index"
	return c.String(http.StatusOK, result)
}

func create(c echo.Context) error {
	title := c.FormValue("title")
	author := c.FormValue("author")

	result := &ResultResponse{
		Message: "You are reached to Create - Crud" + " / " +
			string(title) + " / " +
			string(author),
	}
	return c.JSON(http.StatusOK, result)
}

func read(c echo.Context) error {
	id := c.Param("id")

	result := &ResultResponse{
		Message: "You are reached to Read - cRud" + " / " +
			string(id),
	}

	return c.JSON(http.StatusOK, result)
}

func update(c echo.Context) error {
	resultMap := echo.Map{}
	if err := c.Bind(&resultMap); err != nil {
		return err
	}

	resultMap["Message"] = "You are reached to Update - crUd"

	return c.JSON(http.StatusOK, resultMap)
}

func delete(c echo.Context) error {
	auth := echo.Map{}
	if err := c.Bind(&auth); err != nil {
		return err
	}

	id := c.Param("id")
	result := "You are reached to Delete - cruD" + " / " +
		id + " / " +
		auth["user"].(string) + " / " +
		auth["token"].(string) + " / "
	return c.String(http.StatusOK, result)
}

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		errorResult := &ResultResponse{
			Message: "Contents not found",
		}
		return c.JSON(http.StatusNotFound, errorResult)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/books", index)

	e.POST("/books", create)
	e.GET("/books/:id", read)
	e.PUT("/books", update)
	e.DELETE("/books/:id", delete)

	e.Logger.Fatal(e.Start(":1323"))
}
