package generator

import (
	"sort"

	"github.com/gobuffalo/packr"
	"github.com/roger-russel/goverage/internal/reader"
	"github.com/roger-russel/goverage/internal/template"
	"github.com/spf13/cobra"
)

//BeautifulReport takes the atomic cover profile and make a beautiful html report
func BeautifulReport(c *cobra.Command, args []string, flags map[string]string) {
	coverStruct := reader.ReadFile(flags["coverageFile"])

	filesList := []template.FileList{}

	keys := make([]string, 0, len(coverStruct))
	for k := range coverStruct {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {

		filesList = append(filesList, template.FileList{
			Name:     k,
			Lines:    coverStruct[k].NumberOfStatements,
			Green:    0,
			Red:      0,
			Coverage: 0,
		})

	}

	box := packr.NewBox("../../assets/templates")
	template.Gen(box, filesList)

}
