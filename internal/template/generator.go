package template

import (
	"encoding/json"
	"io"
	"log"
	"text/template"

	"github.com/gobuffalo/packr"
)

//HTMLData is the data struct that will be used on template generator
type HTMLData struct {
	Theme        string
	ThemeBGColor string
	Styles       []string
	Vue          string
	Scripts      []string
	//FileList expect to be a valid json array of fileList
	FilesList    string
	Pages        map[string]*Page
	VisiblePages string
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
	Content string
}

//FileList information to be used on template list of files
type FileList struct {
	Name     string  `json:"name"`
	Lines    int     `json:"lines"`
	Green    int     `json:"green"`
	Red      int     `json:"red"`
	Coverage float32 `json:"coverage"`
}

type VisiblePages struct {
	Current string          `json:"current"`
	List    map[string]bool `json:"list"`
}

//Gen erate Template
func Gen(box packr.Box, wr io.Writer, theme string, filesList []FileList, pages map[string]*Page, visiblePages *VisiblePages) {

	tmpl := getTemplates(box)

	JFileList, err := json.Marshal(filesList)

	if err != nil {
		log.Fatal(err)
	}

	JVisiblePages, err := json.Marshal(visiblePages)

	if err != nil {
		log.Fatal(err)
	}

	HTMLData := HTMLData{

		Theme:        theme,
		ThemeBGColor: getThemeBGColor(theme),

		Styles: []string{
			unBox(box, "css/vuetify.min.css"),
			unBox(box, "css/roboto.css"),
			unBox(box, "css/page.css"),
			unBox(box, "css/table-list.css"),
			unBox(box, "css/themes/"+theme+".css"),
		},

		Scripts: []string{
			unBox(box, "js/vue.min.js"),
			unBox(box, "js/vuetify.min.js"),
		},

		FilesList:    string(JFileList),
		Pages:        pages,
		VisiblePages: string(JVisiblePages),
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
		"table-list.tpl",
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
