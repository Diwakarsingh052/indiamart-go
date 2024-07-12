package db

import (
	"fmt"
)

// bad practice, don't put db conneciton in global exported variables
var Database string = "postgres"

func ReadDB() {
	fmt.Println("reading from ", Database)
}

func GetConn() string {
	return Database
}
