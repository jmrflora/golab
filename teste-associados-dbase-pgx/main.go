package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Valentin-Kaiser/go-dbase/dbase"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Associado struct {
	nome  string
	depro string
}

func main() {
	f, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	dbase.Debug(true, io.MultiWriter(os.Stdout, f))

	table, err := dbase.OpenTable(&dbase.Config{
		Filename:   "CAD01.DBF",
		TrimSpaces: true,
		Untested:   true,
	})
	if err != nil {
		panic(dbase.GetErrorTrace(err))
	}
	defer table.Close()

	for i := 0; i < 2; i++ {
		row, err := table.Next()
		if err != nil {
			panic(dbase.GetErrorTrace(err))
		}

		if row.Deleted {
			fmt.Printf("Deleted row at position: %v \n", row.Position)
			continue
		}
		field_depro := row.Field(0)
		if field_depro == nil {
			panic("field not found")
		}
		field_nome := row.Field(1)
		if field_nome == nil {
			panic("field not found")
		}
		pessoa := new(Associado)
		pessoa.depro = field_depro.GetValue().(string)
		pessoa.nome = field_nome.GetValue().(string)

		println(pessoa.nome)
	}
}
