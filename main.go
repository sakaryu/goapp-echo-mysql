package main

import (
  "fmt"
  "net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type User struct {
  ID int
  Name string
  Token string
}

type Token struct {
  Token string `json:"token"`
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
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
  }))

  e.GET("/api/login", func(c echo.Context) error {
    var user User

    user.ID = 1
    user.Name = "test"
    user.Token = "hotreload test"

    return c.JSON(
      http.StatusOK,
      user,
    )
  })

  e.POST("/api/login", func(c echo.Context) error {
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

    var token Token

    // for rows.Next() {
    //   err := rows.Scan(&user.ID, &user.Name)
    //   if err != nil {
    //     panic(err.Error())
    //   }
    //   fmt.Println(user.ID, user.Name)
    // }

    user.ID = 0
    user.Name = "aaa"
    token.Token = "sample"

    fmt.Println(token)
    fmt.Println(user)

    err = rows.Err()
    if err != nil {
      panic(err.Error())
    }

    return c.JSON(
      http.StatusOK,
      token,
    )
  })

  e.GET("/api/me", func(c echo.Context) error {
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
