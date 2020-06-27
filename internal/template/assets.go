package template

import (
	"bytes"
	"log"
	"text/template"

	"github.com/markbates/pkger"
)

var baseDir string = "/assets/templates"

//getTemplates Reader the templates and pass to template library
func getTemplates() (tmpl *template.Template) {

	f, err := pkger.Open(baseDir + "/index.tpl")

	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(f)
	f.Close()

	tmpl = template.Must(template.New("index").Parse(buf.String()))

	templatesNames := []string{
		"table-list.tpl",
		"page.tpl",
	}

	for _, tplName := range templatesNames {

		f, err := pkger.Open(baseDir + "/" + tplName)

		if err != nil {
			log.Fatal(err)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(f)
		f.Close()

		tpl := buf.String()
		tmpl = template.Must(tmpl.Parse(tpl))

	}

	return tmpl
}

func unBox(path string) string {

	f, err := pkger.Open(baseDir + "/" + path)

	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(f)
	f.Close()

	return buf.String()

}
