package main

//	"github.com/gobuffalo/packr"
//  "github.com/roger-russel/goverage/internal/template"

import (
	"fmt"

	"github.com/roger-russel/goverage/internal/cmd"
)

var version string
var commit string
var date string

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some thing went wrong:", r)
		}
	}()

	cmd.Execute(cmd.FullVersion{
		Version: version,
		Commit:  commit,
		Date:    date,
	})

}
