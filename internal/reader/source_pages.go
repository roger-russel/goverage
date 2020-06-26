package reader

import (
	"sort"
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

func parsePage(wg *sync.WaitGroup, chn chan template.Page, path string, name string, rawContent *LineCover) {
	/*
		defer func() {
			if r := recover(); r != nil {
				log.Println("fail parsing", name, " error:", r)
				(*wg).Done()
			}
		}()
	*/
	var page template.Page

	page.FullName = name

	if path[len(path)-1:] != "/" {
		path += "/"
	}

	fileScanner, rd := read(path + name)
	defer rd.Close()

	var lnNumber int
	var coveredCount int
	var tracked bool

	for fileScanner.Scan() {
		lnNumber++

		line := template.Line{
			Line:     lnNumber,
			Contents: []template.Content{},
		}

		var content string = fileScanner.Text()
		var lineContent template.Content

		if columns, ok := (*rawContent).Report.CountThroughtLines[lnNumber]; ok {

			keysColumn := []int{}
			for k := range columns {
				keysColumn = append(keysColumn, k)
			}

			sort.Ints(keysColumn)

			var previousColumn int = 0
			var previousSliceColumn int = 0

			for _, colNumber := range keysColumn {

				ccNumber := (*rawContent).Report.CountThroughtLines[lnNumber][colNumber]

				tracked = verifyIfTracked(&(*rawContent).Report.Tracked, lnNumber, previousColumn)

				left := content[previousSliceColumn:(colNumber - 1)]

				previousColumn = colNumber
				previousSliceColumn = previousColumn - 1

				lineContent = template.Content{
					Tracked: tracked,
					Count:   coveredCount,
					Content: left,
				}

				line.Contents = append(line.Contents, lineContent)

				coveredCount += ccNumber

			}

			if previousSliceColumn != len(content) {

				tracked = verifyIfTracked(&(*rawContent).Report.Tracked, lnNumber, previousColumn)

				lineContent = template.Content{
					Tracked: tracked,
					Count:   coveredCount,
					Content: content[previousSliceColumn-1:],
				}

				line.Contents = append(line.Contents, lineContent)

			}
		} else {

			tracked = verifyIfTracked(&(*rawContent).Report.Tracked, lnNumber, 0)

			lineContent = template.Content{
				Tracked: tracked,
				Count:   coveredCount,
				Content: content,
			}

			line.Contents = append(line.Contents, lineContent)

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

func verifyIfTracked(trackRange *[]Tracked, line int, col int) bool {

	for _, r := range *trackRange {

		if r.StartLine <= line &&
			r.EndLine >= line {

			if (r.StartLine == line && r.StartCol > col) ||
				(r.EndLine == line && r.EndCol < col) {
				continue
			}

			return true

		}

	}

	return false

}
