package main

//	"github.com/gobuffalo/packr"
//  "github.com/roger-russel/goverage/internal/template"

import "github.com/roger-russel/goverage/internal/cmd"

var version string
var commit string
var date string

//FullVersion is a struct with version information

func main() {

	//	box := packr.NewBox("../../assets/templates")
	// template.Gen(box)

	cmd.Execute(cmd.FullVersion{
		Version: version,
		Commit:  commit,
		Date:    date,
	})

}
