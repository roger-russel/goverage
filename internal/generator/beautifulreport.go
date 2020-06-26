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

	coverStruct := reader.CoverageReport(flags.CoverageFile)

	box := packr.NewBox("../../assets/templates")

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

	//pages = make(map[string]*template.Page, 0)
	//fmt.Println("pages", len(pages), render.AsCode(pages))
	/*
		pagetmp := []template.Page{
			{
				FullName: "./cmp/main.go",
				Lines: []template.Line{
					{
						Line: 1,
						Contents: []template.Content{
							{
								Tracked: false,
								Count:   0,
								Content: "package main",
							},
						},
					},
					{
						Line: 2,
						Contents: []template.Content{
							{
								Tracked: false,
								Count:   0,
								Content: "var b int",
							},
						},
					},
					{
						Line: 3,
						Contents: []template.Content{
							{
								Tracked: true,
								Count:   3,
								Content: "func main() {",
							},
						},
					},
					{
						Line: 4,
						Contents: []template.Content{
							{
								Tracked: true,
								Count:   4,
								Content: html.EscapeString("  if a := 1;"),
							},
							{
								Tracked: true,
								Count:   4,
								Content: html.EscapeString("a < b;"),
							},
							{
								Tracked: true,
								Count:   0,
								Content: html.EscapeString("a > b"),
							},
							{
								Tracked: false,
								Count:   0,
								Content: "{",
							},
						},
					},
					{
						Line: 5,
						Contents: []template.Content{
							{
								Tracked: true,
								Count:   0,
								Content: html.EscapeString("    fmt.Println(\"a<b\")"),
							},
						},
					},
					{
						Line: 6,
						Contents: []template.Content{
							{
								Tracked: false,
								Count:   2,
								Content: "  }",
							},
						},
					},
					{
						Line: 7,
						Contents: []template.Content{
							{
								Tracked: false,
								Count:   0,
								Content: "}",
							},
						},
					},
				},
			},
		}
	*/
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

	template.Gen(box, wr, flags.Theme, filesList, pages, &visiblePages)

}
