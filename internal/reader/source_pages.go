package reader

import (
	"log"
	"sync"

	"github.com/roger-russel/goverage/internal/template"
)

//SourcePages reade the file source to generate the page coverage struct report
func SourcePages(rootPath string, coverStruct *CoverStruct) map[string]*template.Page {

	pages := make(map[string]*template.Page)
	ch := make(chan template.Page, 1)

	var wg sync.WaitGroup
	pagesNumber := len((*coverStruct))

	wg.Add(pagesNumber)

	for name, content := range *coverStruct {
		go parsePage(&wg, ch, rootPath, name, content)
	}

	go chnPage(&wg, ch, &pages)

	wg.Wait()

	return pages

}

func parsePage(wg *sync.WaitGroup, chn chan template.Page, path string, name string, content *LineCover) {

	defer func() {
		if r := recover(); r != nil {
			log.Println("fail parsing", name, " error:", r)
			(*wg).Done()
		}
	}()

	var page template.Page

	page.FullName = name

	if path[len(path)-1:] != "/" {
		path += "/"
	}

	fileScanner, rd := read(path + name)
	defer rd.Close()

	var lnNumber int
	for fileScanner.Scan() {
		lnNumber++
		line := template.Line{
			Line: lnNumber,
			Contents: []template.Content{
				{
					Tracked: false,
					Count:   0,
					Content: fileScanner.Text(),
				},
			},
		}
		page.Lines = append(page.Lines, line)
	}

	chn <- page

}

func chnPage(wg *sync.WaitGroup, chn chan template.Page, pages *map[string]*template.Page) {

	for p := range chn {
		page := p
		(*pages)[p.FullName] = &page
		wg.Done()
	}

}
