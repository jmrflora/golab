package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Valentin-Kaiser/go-dbase/dbase"
)

func main() {
	f, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
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
	defer table.Close()

	fmt.Printf(
		"Last modified: %v Columns count: %v Record count: %v File size: %v \n",
		table.Header().Modified(0),
		table.Header().ColumnsCount(),
		table.Header().RecordsCount(),
		table.Header().FileSize(),
	)
	for _, column := range table.Columns() {
		fmt.Printf("Name: %v - Type: %v \n", column.Name(), column.Type())
	}

	// Init the field we want to search for.
	// Search for a product containing the word "test" in the name.

	// field, err := table.NewField(1, "ADAILVA")
	// if err != nil {
	// 	panic(dbase.GetErrorTrace(err))
	// }

	// records, err := table.Search(field, false)
	// if err != nil {
	// 	panic(dbase.GetErrorTrace(err))
	// }

	// // Print all found records.
	// fmt.Println("Found records without exact match:")
	// for _, record := range records {

	// 	for _, field := range record.Fields() {

	// 		fmt.Printf("%v %v \n", field.Name(), field.GetValue())

	// 	}

	// 	// field = record.Field(1)
	// 	// if field == nil {
	// 	// 	panic("Field 'NOMEx' not found")
	// 	// }

	// 	// fmt.Printf("%v \n", field.GetValue())
	// }

	// !table.EOF()
	for i := 0; i < 10; i++ {
		row, err := table.Next()
		if err != nil {
			panic(dbase.GetErrorTrace(err))
		}

		// Skip deleted rows.
		if row.Deleted {
			fmt.Printf("Deleted row at position: %v \n", row.Position)
			continue
		}

		// Get the first field by column position
		field := row.Field(1)
		if field == nil {
			panic("Field not found")
		}

		// Print the field value.
		fmt.Printf("Field: %v [%v] => %v \n", field.Name(), field.Type(), field.GetValue())
	}

}
