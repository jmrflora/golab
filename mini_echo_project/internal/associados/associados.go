package associados

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Associado struct {
	Id       string
	Nome     string `json:"nome" validate:"required"`
	Endereco string `json:"endereco" validate:"required"`
}

type CustomValidator struct {
	Validator *validator.Validate
}

type fake_db struct {
	Associados map[int]*Associado
	Seq        int
}

func (fdb *fake_db) Append(A *Associado) {
	i, err := strconv.Atoi(A.Id)
	if err != nil {
		panic(err)
	}
	fdb.Associados[i] = A
	// println(fdb.Associados[i].Nome)
	fdb.Seq++

}

func NewDb() *fake_db {
	return &fake_db{
		Associados: make(map[int]*Associado),
		Seq:        1,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
