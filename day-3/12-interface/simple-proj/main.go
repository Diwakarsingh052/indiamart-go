package main

import (
	"simple-proj/stores"
	"simple-proj/stores/mysql"
	"simple-proj/stores/postgres"
)

func main() {
	u := stores.User{
		Name:  "ajay",
		Email: "ajay@email.com",
	}
	m := mysql.New("mysql")
	p := postgres.New("postgres")
	ms := stores.NewService(m)
	ms.Create(u)
	ps := stores.NewService(p)
	ps.Create(u)
}
