package main

import "log"

type Conf struct {
	db string // unexported field
}

// NewConf function is created to create instance of the Conf struct
// any external package can use NewConf to have a db connection
// note db string is unexported, which means no one from the outside package could change it
func NewConf(dataSourceName string) *Conf {

	if dataSourceName == "" {
		log.Println("please provide datasource name")
		return nil
	}

	return &Conf{db: dataSourceName}
}
