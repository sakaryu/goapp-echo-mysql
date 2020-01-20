package main

import (
  "fmt"
  "net/http"
  "github.com/labstack/echo"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type User struct {
  ID int
  Name string
}

func main() {
  db, err := sql.Open("mysql", "root:password@tcp(godockerDB)/sample_dev")
  if err != nil {
    panic(err)
  }

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  defer db.Close()

  e := echo.New()
  e.GET("/sample", func(c echo.Context) error {
    db, err := sql.Open("mysql", "root:password@tcp(godockerDB)/sample_dev")
    if err != nil {
      panic(err)
    }

    err = db.Ping()
    if err != nil {
      panic(err)
    }

    defer db.Close()

    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
      panic(err.Error())
    }
    defer rows.Close()

    var user User

    for rows.Next() {
      err := rows.Scan(&user.ID, &user.Name)
      if err != nil {
        panic(err.Error())
      }
      fmt.Println(user.ID, user.Name)
    }

    err = rows.Err()
    if err != nil {
      panic(err.Error())
    }

    return c.JSON(
      http.StatusOK,
      user,
    )
  })
  e.Logger.Fatal(e.Start(":8080"))
}
