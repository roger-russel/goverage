package template

import (
	"encoding/json"
	"html"
	"io"
	"log"
	"text/template"

	"github.com/gobuffalo/packr"
)

//HTMLData is the data struct that will be used on template generator
type HTMLData struct {
	Theme   string
	Styles  []string
	Vue     string
	Scripts []string
	//FileList expect to be a valid json array of fileList
	FilesList string
	Pages     []Page
}

//Page Information
type Page struct {
	FullName string
	Lines    []Line
}

//Line information
type Line struct {
	Line     int
	Contents []Content
}

//Content of lines
type Content struct {
	Tracked bool // If the content is tracked by coverage or not
	Count   int
	Color   string
	Content string
}

//FileList information to be used on template list of files
type FileList struct {
	Name     string  `json:"name"`
	Lines    int     `json:"lines"`
	Coverage float32 `json:"coverage"`
}

//Gen erate Template
func Gen(box packr.Box, filesList []FileList, wr io.Writer, theme string) {

	tmpl := getTemplates(box)

	JFileList, err := json.Marshal(filesList)

	if err != nil {
		log.Fatal(err)
	}

	HTMLData := HTMLData{

		Theme: theme,

		Styles: []string{
			unBox(box, "css/vuetify.min.css"),
			unBox(box, "css/roboto.css"),
			unBox(box, "css/page.css"),
			unBox(box, "css/themes/"+theme+".css"),
		},

		Scripts: []string{
			unBox(box, "js/vue.min.js"),
			unBox(box, "js/vuetify.min.js"),
		},

		FilesList: string(JFileList),
		Pages: []Page{
			{
				FullName: "./cmp/main.go",
				Lines: []Line{
					{
						Line: 1,
						Contents: []Content{
							{
								Tracked: false,
								Count:   0,
								Content: "package main",
							},
						},
					},
					{
						Line: 2,
						Contents: []Content{
							{
								Tracked: false,
								Count:   0,
								Content: "var b int",
							},
						},
					},
					{
						Line: 3,
						Contents: []Content{
							{
								Tracked: true,
								Count:   3,
								Content: "func main() {",
							},
						},
					},
					{
						Line: 4,
						Contents: []Content{
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
						Contents: []Content{
							{
								Tracked: true,
								Count:   0,
								Content: html.EscapeString("    fmt.Println(\"a<b\")"),
							},
						},
					},
					{
						Line: 6,
						Contents: []Content{
							{
								Tracked: false,
								Count:   2,
								Content: "  }",
							},
						},
					},
					{
						Line: 7,
						Contents: []Content{
							{
								Tracked: false,
								Count:   0,
								Content: "}",
							},
						},
					},
				},
			},
		},
	}

	tmpl.Execute(wr, HTMLData)

}

func unBox(box packr.Box, path string) string {
	pack, err := box.FindString(path)

	if err != nil {
		log.Fatal(err)
	}

	return pack
}

func getTemplates(box packr.Box) (tmpl *template.Template) {

	tplIndex, err := box.FindString("index.tpl")

	if err != nil {
		log.Fatal(err)
	}

	tmpl = template.Must(template.New("index").Parse(tplIndex))

	templatesNames := []string{
		"page.tpl",
	}

	for _, tplName := range templatesNames {

		tpl, err := box.FindString(tplName)

		if err != nil {
			log.Fatal(err)
		}

		tmpl = template.Must(tmpl.Parse(tpl))

	}

	return tmpl

}
