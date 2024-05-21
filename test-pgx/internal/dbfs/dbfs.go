package dbfs

import (
	"fmt"
	"io"
	"os"

	"github.com/Valentin-Kaiser/go-dbase/dbase"
)

func RetornarTabela() (*dbase.File, error) {
	f, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	dbase.Debug(true, io.MultiWriter(os.Stdout, f))

	// Open the example database table.
	table, err := dbase.OpenTable(&dbase.Config{
		Filename:   "CAD01.DBF",
		TrimSpaces: true,
		Untested:   true,
	})
	if err != nil {
		panic(dbase.GetErrorTrace(err))
	}
	return table, nil
}
