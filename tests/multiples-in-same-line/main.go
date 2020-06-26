package main

import (
	"fmt"
)

func main() {

	if a, b, c := getMultiplesResponses(); a && b && c {
		fmt.Println("In")
	}

	if err := getError(); err != nil { fmt.Println(err) }

	if true == true { if false != false  { fmt.Println("false") } else { fmt.Println("true") } }
}

func getMultiplesResponses() (bool, bool, bool) {

	return true, true, false

}

func getError() error {
	return fmt.Errorf("error")
}
