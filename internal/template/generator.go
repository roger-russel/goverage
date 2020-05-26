package template

import (
	"encoding/json"
	"log"
	"os"
	"text/template"

	"github.com/gobuffalo/packr"
)

//HTMLData is the data struct that will be used on template generator
type HTMLData struct {
	Styles  []string
	Vue     string
	Scripts []string
	//FileList expect to be a valid json array of fileList
	FilesList string
}

//FileList information to be used on template list of files
type FileList struct {
	Name     string  `json:"name"`
	Lines    int     `json:"lines"`
	Green    int     `json:"green"`
	Red      int     `json:"red"`
	Coverage float32 `json:"coverage"`
}

//Gen erate Template
func Gen(box packr.Box, filesList []FileList) {

	htmlTPL, err := box.FindString("index.tpl")

	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("index").Parse(htmlTPL)

	if err != nil {
		log.Fatal(err)
	}

	JFileList, err := json.Marshal(filesList)

	if err != nil {
		log.Fatal(err)
	}

	HTMLData := HTMLData{
		Styles: []string{
			unBox(box, "css/vuetify.min.css"),
			unBox(box, "css/roboto.css"),
		},

		Scripts: []string{
			unBox(box, "js/vue.min.js"),
			unBox(box, "js/vuetify.min.js"),
		},

		FilesList: string(JFileList),
	}

	err = tmpl.Execute(os.Stdout, HTMLData)

	if err != nil {
		panic(err)
	}

}

func unBox(box packr.Box, path string) string {
	pack, err := box.FindString(path)

	if err != nil {
		log.Fatal(err)
	}

	return pack
}
