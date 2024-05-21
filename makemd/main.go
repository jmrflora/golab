package main

import "os"

func main() {
	file, err := os.Create("tstemd")
	if err != nil {
		println(err.Error())
		return
	}
	defer file.Close()

	_, err = file.WriteString("# titulo\n\n## sub\n\naaaaa\n\naaaaa\n* 1\n* 2\n* 2\n\n222")
	if err != nil {
		println(err.Error())
		return
	}
}
