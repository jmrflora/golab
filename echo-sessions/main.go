package main

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// var store *sqlitestore.SqliteStore

type Person struct {
	Id   int
	Nome string
}

type M map[string]interface{}

func init() {

	gob.Register(&Person{})
	gob.Register(&M{})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("super-secret-key"))))

	e.GET("/", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   60 * 60,
			HttpOnly: true,
		}
		p := &Person{
			Id:   10,
			Nome: "jose",
		}
		sess.Values["person"] = p
		sess.Save(c.Request(), c.Response())
		return c.NoContent(http.StatusOK)
	})

	e.GET("/try", func(c echo.Context) error {

		sess, err := session.Get("session", c)
		if err != nil {
			return echo.NewHTTPError(http.StatusMethodNotAllowed, err.Error())
		}
		val := sess.Values["person"]

		if person, ok := val.(*Person); !ok {
			// println(person.Nome)
			return echo.ErrBadRequest
		} else {

			return c.JSON(http.StatusOK, person)
		}
	})

	e.Logger.Fatal(e.Start(":1323"))
}
