package template

import (
	"log"
	"os"
	"text/template"

	"github.com/gobuffalo/packr"
)

type HtmlData struct {
	Styles  []string
	Vue     string
	Scripts []string
}

//Gen erate Template
func Gen(box packr.Box) {

	htmlTPL, err := box.FindString("index.tpl")

	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("index").Parse(htmlTPL)

	if err != nil {
		log.Fatal(err)
	}

	htmlData := HtmlData{
		Styles: []string{
			unBox(box, "css/vuetify.min.css"),
		},
		Scripts: []string{
			unBox(box, "js/vue.min.js"),
			unBox(box, "js/vuex.js"),
			unBox(box, "js/vuetify.min.js"),
		},
	}

	err = tmpl.Execute(os.Stdout, htmlData)

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
