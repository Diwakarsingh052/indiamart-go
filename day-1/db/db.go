package db

import "fmt"

var Database string = "postgres"

func ReadDB() {
	fmt.Println("reading from ", Database)
}

func GetConn() string {
	return Database
}
