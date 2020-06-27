package generator

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/roger-russel/goverage/internal/flags"
	"github.com/roger-russel/goverage/internal/reader"
	"github.com/roger-russel/goverage/internal/template"
	"github.com/spf13/cobra"
)

//BeautifulReport takes the atomic cover profile and make a beautiful html report
func BeautifulReport(c *cobra.Command, args []string, flags flags.Flags) {

	coverStruct := reader.CoverageReport(flags.CoverageFile)

	filesList := []template.FileList{}

	keys := make([]string, 0, len(coverStruct))

	var visiblePages template.VisiblePages
	visiblePages.List = make(map[string]bool, 0)

	for k := range coverStruct {
		keys = append(keys, k)
		visiblePages.List[k] = false
	}

	sort.Strings(keys)

	for _, k := range keys {

		filesList = append(filesList, template.FileList{
			Name:     k,
			Lines:    coverStruct[k].NumberOfStatements,
			Green:    coverStruct[k].Green,
			Red:      coverStruct[k].Red,
			Coverage: coverStruct[k].Coverage,
		})

	}

	pages := reader.SourcePages(flags.Path, &coverStruct)

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

	template.Gen(wr, flags.Theme, filesList, pages, &visiblePages)

}
