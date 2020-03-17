package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/go-playground/validator/v10"
)

// User type
type User struct {
  ID int64 `json:"id"`
  Name string `json:"name" validate:"required"`
  Email string `json:"email" validate:"required,email"`
}

// Validator type
type Validator struct {
  validator *validator.Validate
}

// Validate validate
func (v *Validator) Validate(i interface{}) error {
  return v.validator.Struct(i)
}

func main() {
  db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/sample_dev")
  if err != nil {
    panic(err)
  }

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  defer db.Close()

  e := echo.New()
  e.Validator = &Validator{validator: validator.New()}
  e.Static("/", "static/")

  e.GET("/api/users", func(c echo.Context) error {
    db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/sample_dev")
    if err != nil {
      panic(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
      panic(err.Error())
    }
    defer rows.Close()

    var users []User
    var user User

    for rows.Next() {
      err := rows.Scan(&user.ID, &user.Name, &user.Email)
      if err != nil {
        panic(err.Error())
      }
      users = append(users, user)
    }

    return c.JSON(
      http.StatusOK,
      users,
    )
  })

  e.GET("/api/users/:id", func(c echo.Context) error {
    db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/sample_dev")
    if err != nil {
      panic(err)
    }
    defer db.Close()

    var user User

    row := db.QueryRow(`SELECT * FROM users WHERE id = ? LIMIT 1`, c.Param("id"))

    err = row.Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        c.Logger().Error("Select: ", err)
        return c.String(
          http.StatusBadRequest,
          "Select: "+err.Error(),
        )
    }

    return c.JSON(
      http.StatusOK,
      user,
    )
  })

  e.POST("/api/users", func(c echo.Context) error {
    db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/sample_dev")
    if err != nil {
      panic(err)
    }
    defer db.Close()

    var user User

		if err = c.Bind(&user); err != nil {
      c.Logger().Error("Bind: ", err)
      return c.String(
        http.StatusBadRequest,
        "Bind: "+err.Error(),
      )
		}

		if err = c.Validate(&user); err != nil {
      c.Logger().Error("Validate: ", err)
      return c.String(
        http.StatusBadRequest,
        "Validate: "+err.Error(),
      )
		}

    _, err = db.Exec(`INSERT INTO users(name, email) VALUES(?, ?)`, user.Name, user.Email)
    if err != nil {
      c.Logger().Error("Insert: ", err)
      return c.String(
        http.StatusBadRequest,
        "Insert: "+err.Error(),
      )
    }

    return c.JSON(
      http.StatusCreated,
      "",
    )
  })
  e.Logger.Fatal(e.Start(":8080"))
}
