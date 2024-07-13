package main

import (
	"errors"
	"fmt"
)

//Error should be suffixed in the name of the struct use for error handling

type QueryError struct {
	Func  string
	Input string
	Err   error
}

// implementing the error interface // format your msg here
func (q *QueryError) Error() string {
	return "main." + q.Func + ": " + "input " + q.Input + " " + q.Err.Error()
}

func main() {
	//fmt.Println(strconv.Atoi("xyz"))
	//fmt.Println(strconv.Atoi("abc"))
	//fmt.Println(strconv.ParseInt("efg", 10, 64))
	//os.PathError{}

	err := SearchSomething("abc")
	if err != nil {
		fmt.Println(err)
		var qe *QueryError
		// errors.As check if a struct is present in the chain or not for error,
		// if yes than fill the values in qe var with the values of QueryError type
		if errors.As(err, &qe) {
			fmt.Println("QueryError:", qe.Func)
			return
		}
		return
	}

}

func SearchSomething(s string) error {
	//this is the case when data is not present
	return &QueryError{
		Func:  "SearchSomething",
		Input: s,
		Err:   errors.New("not found"),
	}
}
