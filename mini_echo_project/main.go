package main

import (
	"example/echo-project/internal/associados"
	"example/echo-project/internal/templates"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/a-h/templ"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var fdb = associados.NewDb()
var lock = sync.Mutex{}

func main() {
	e := echo.New()
	e.Validator = &associados.CustomValidator{Validator: validator.New()}

	e.POST("/associados", CriarAssociado)
	e.GET("/associados/:id", LerAssociado)
	e.GET("/teste", Teste)
	e.GET("/click", click)
	e.Logger.Fatal(e.Start(":1323"))
}

func CriarAssociado(c echo.Context) (err error) {
	lock.Lock()
	defer lock.Unlock()
	a := &associados.Associado{
		Id: fmt.Sprint(fdb.Seq),
	}
	if err = c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(a); err != nil {
		return err
	}
	fdb.Append(a)
	return c.JSON(http.StatusOK, a)

}

func LerAssociado(c echo.Context) (err error) {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	// println(id)
	// println(fdb.Associados[id].Nome)
	return c.JSON(http.StatusOK, fdb.Associados[id])
}

func Teste(c echo.Context) (err error) {
	lock.Lock()
	defer lock.Unlock()
	cmp := templates.Hello("ola")
	return render(c, 200, cmp)
}
func render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)

	err := t.Render(ctx.Request().Context(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil
}

func click(c echo.Context) (err error) {
	lock.Lock()
	defer lock.Unlock()
	println("aquiiii")
	return c.String(http.StatusOK, "Hello, World!")
}
