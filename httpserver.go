package main

import (
	"database/sql"
	"log"
	"net/http"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type testS struct {
	Name string
	Fuck string
}

func getUser(c echo.Context) error {
	// User ID 来自于url `users/:id`
	// id := c.Param("id")
	db, err := sql.Open("mysql", "root:123456@/mysql")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM TEST")
	if err != nil {
		log.Fatal(err)
	}

	var tests = []testS{}
	for rows.Next() {
		var test testS
		rows.Scan(&test.Name, &test.Fuck)
		tests = append(tests, test)
	}
	//fmt.Println(test.Name, test.Fuck)
	defer db.Close()
	defer rows.Close()
	b, err := json.Marshal(tests)
	if err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, "Hello, World!:"+"\n"+string(b))
}

func main() {
	e := echo.New()
	e.Use(middleware.Gzip())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)

	e.Logger.Fatal(e.Start(":8081"))
}
