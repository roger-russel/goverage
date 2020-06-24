package generator

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/gobuffalo/packr"
	"github.com/roger-russel/goverage/internal/flags"
	"github.com/roger-russel/goverage/internal/reader"
	"github.com/roger-russel/goverage/internal/template"
	"github.com/spf13/cobra"
)

//BeautifulReport takes the atomic cover profile and make a beautiful html report
func BeautifulReport(c *cobra.Command, args []string, flags flags.Flags) {

	coverStruct := reader.ReadFile(flags.CoverageFile)

	box := packr.NewBox("../../assets/templates")

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

	var wr io.Writer

	if output := flags.Output; output != "" {

		f, err := os.Create(output)

		if err != nil {
			panic(fmt.Sprintf("Could not write into %v, error: %v", output, err))
		}

		wr = f
		defer f.Close()

	} else {
		wr = os.Stdout
	}

	template.Gen(box, filesList, wr, flags.Theme)

}
