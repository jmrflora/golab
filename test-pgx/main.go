package main

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Pessoa struct {
	Id    int
	Name  string
	Email string
}

func main() {
	db, err := sqlx.Open("pgx", "postgresql://jmrflora:100231@localhost:5432/teste_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	p := Pessoa{
		Id:    1,
		Name:  "jonh doe",
		Email: "jonhdoe@email.com",
	}
	pp := []Pessoa{}

	rows, err := db.NamedQuery("select * from users where id=:id", p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	for rows.Next() {
		var p2 Pessoa
		err = rows.StructScan(&p2)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(p2.Name)
	}

	rows, err = db.Queryx("select * from users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	i := 0
	for rows.Next() {
		var p0 Pessoa

		err = rows.StructScan(&p0)
		pp = append(pp, p0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(i)
		fmt.Println(pp[i].Name)
		i++
	}
	// for i, v := range pp {
	// 	fmt.Printf("Index: %d, nome: %v\n", i, v.Name)
	// }

}
