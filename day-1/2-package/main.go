package main

import (
	"fmt"
	"github.com/username/reponame/day-1/db"
	"github.com/username/reponame/day-1/sum"
)

func main() {
	sum.Add()
	db.Database = "abc"
	fmt.Println(db.GetConn())
}
