package main

//	"github.com/gobuffalo/packr"
//  "github.com/roger-russel/goverage/internal/template"

import "github.com/roger-russel/goverage/internal/cmd"

var version string
var commit string
var date string

func main() {

	cmd.Execute(cmd.FullVersion{
		Version: version,
		Commit:  commit,
		Date:    date,
	})

}
