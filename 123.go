package main

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	)

type qwerty struct{
	id int
	Name string

}
func main() {

	connStr := "user=sa password=sa dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from t_itfb")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var Name []qwerty

	for rows.Next() {
		p := qwerty{}
		err := rows.Scan(&p.id, &p.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Name = append(Name, p)
	}
	//Запихнул все в переменную Name
	for _, p := range Name {
		fmt.Println(p.id, p.Name)
	}
	//Открываем второй конект к второй базе
	connStr2 := "user=sa password=sa dbname=testdb2 sslmode=disable"
	db2, err := sql.Open("postgres", connStr2)
	if err != nil {
		panic(err)
	}
	defer db2.Close()

	//Передаю переменную Name в цикл, и пробегаю инстертом по полученным строкам
	for _, p := range Name {
		fmt.Println(p.id, p.Name)
		result, err := db2.Exec("insert into public.t_itfb(id, name) VALUES ($1, $2)",
			p.id, p.Name)
		if err != nil{
			panic(err)
			continue
		}

		fmt.Println(result.RowsAffected())
	}

}