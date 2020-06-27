package template

import (
	"encoding/json"
	"io"
	"log"
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
	PagesInJSON  string
}

//Page Information
type Page struct {
	FullName string `json:"full-name"`
	Lines    []Line `json:"lines"`
}

//Line information
type Line struct {
	Line     int       `json:"line"`
	Contents []Content `json:"contents"`
}

//Content of lines
type Content struct {
	Tracked bool   `json:"tracked"` // If the content is tracked by coverage or not
	Count   int    `json:"count"`
	Content string `json:"content"`
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
func Gen(wr io.Writer, theme string, filesList []FileList, pages map[string]*Page, visiblePages *VisiblePages) {

	tmpl := getTemplates()

	JFileList, err := json.Marshal(filesList)

	if err != nil {
		log.Fatal(err)
	}

	JVisiblePages, err := json.Marshal(visiblePages)

	if err != nil {
		log.Fatal(err)
	}

	jPage, err := json.Marshal(pages)

	if err != nil {
		log.Fatal(err)
	}

	HTMLData := HTMLData{

		Theme:        theme,
		ThemeBGColor: getThemeBGColor(theme),

		Styles: []string{
			unBox("css/vuetify.min.css"),
			unBox("css/roboto.css"),
			unBox("css/page.css"),
			unBox("css/table-list.css"),
			unBox("css/themes/" + theme + ".css"),
		},

		Scripts: []string{
			unBox("js/vue.min.js"),
			unBox("js/vuetify.min.js"),
		},

		FilesList:    string(JFileList),
		Pages:        pages,
		VisiblePages: string(JVisiblePages),
		PagesInJSON:  string(jPage),
	}

	tmpl.Execute(wr, HTMLData)

}
