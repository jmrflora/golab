package main

import "github.com/LindsayBradford/go-dbf/godbf"

func main() {
	dbfTable, err := godbf.NewFromFile("CAD01.DBF", "iso-8859-1")
	if err != nil {
		println("ola")
		println(err.Error())
	}
	println(dbfTable.NumberOfRecords())

	// exampleList := make(ExampleList, dbfTable.NumberOfRecords())

	// for i := 0; i < dbfTable.NumberOfRecords(); i++ {
	// 	// exampleList[i] = new(ExampleListEntry)
	// 	var s1 string
	// 	s1, err = dbfTable.FieldValueByName(i, "NOMEx")
	// 	if err != nil {
	// 		println("bbbbbbbbbbbbbbbb")
	// 	}
	// 	println(s1)
	// }
}
