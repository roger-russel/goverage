package reader

import (
	"strconv"
	"strings"
)

//LineCover is the page coverage data
type LineCover struct {
	//NumberOfStatements in line
	Coverage           float32
	NumberOfStatements int
	Green              int
	Red                int

	Report Report
}

type Report struct {
	//        line   column count
	CountThroughtLines map[int]map[int]int
	Tracked            []Tracked
}

type Tracked struct {
	StartLine int
	StartCol  int
	EndLine   int
	EndCol    int
}

//CoverStruct       filename
type CoverStruct map[string]*LineCover

type coverCount struct {
	Statements int
	Cover      bool
}

//CoverageReport read an atomic coverage file and turn it into a map
func CoverageReport(filename string) CoverStruct {

	fileScanner, reader := read(filename)
	defer reader.Close()

	if fileScanner.Scan() {
		if strings.Trim(fileScanner.Text(), "") != "mode: atomic" {
			panic("only model atomic is supported")
		}
	}

	coverStruct := make(CoverStruct, 0)
	duplicatedHash := make(map[string]map[string]bool)
	cCont := make(map[string]map[string]coverCount)

	for fileScanner.Scan() {
		swp := fileScanner.Text()
		if swp == "" {
			continue
		}
		line := strings.Split(swp, ":")
		file := line[0]
		splitContent(coverStruct, file, line[1], duplicatedHash, cCont)
	}

	for page := range cCont {

		pgCountStatments := 0
		pgCountCover := 0

		for _, cc := range cCont[page] {
			pgCountStatments += cc.Statements
			if cc.Cover {
				pgCountCover += cc.Statements
			}
		}

		(*coverStruct[page]).Green = pgCountCover
		(*coverStruct[page]).Red = pgCountStatments - pgCountCover
		(*coverStruct[page]).Coverage = float32(pgCountCover) * 100 / float32(pgCountStatments)

	}

	return coverStruct

}

func splitContent(c CoverStruct, file string, lineRaw string, duplicatedHash map[string]map[string]bool, cCount map[string]map[string]coverCount) {

	swapRaw := strings.Split(lineRaw, " ")
	count, err := strconv.Atoi(swapRaw[2])
	area := swapRaw[0]

	if err != nil {
		panic(err)
	}

	if _, ok := c[file]; !ok {
		c[file] = &LineCover{
			NumberOfStatements: 0,
			Coverage:           0,
			Green:              0,
			Red:                0,
			Report: Report{
				CountThroughtLines: make(map[int]map[int]int, 0),
				Tracked:            make([]Tracked, 0),
			},
		}
	}

	statements, err := strconv.Atoi(swapRaw[1])

	if err != nil {
		panic(err)
	}

	if _, ok := duplicatedHash[file]; !ok {
		duplicatedHash[file] = make(map[string]bool)
	}

	if isDuplicated, ok := duplicatedHash[file][area]; !(isDuplicated || ok) {
		duplicatedHash[file][area] = true
		c[file].NumberOfStatements += statements
	}

	if _, ok := cCount[file]; !ok {
		cCount[file] = make(map[string]coverCount)
	}

	if cc, ok := cCount[file][area]; !cc.Cover || !ok {
		cCount[file][area] = coverCount{
			Statements: statements,
			Cover:      count > 0,
		}
	}

	swapRawLineCol := strings.Split(area, ",")

	startLine, startCol := splitLineAndCol(swapRawLineCol[0])
	endLine, endCol := splitLineAndCol(swapRawLineCol[1])

	addContent(c, file, startLine, startCol, count)
	//Because it is end of it will not increase the statement on new lines
	addContent(c, file, endLine, endCol, count*-1)

	c[file].Report.Tracked = append(c[file].Report.Tracked, Tracked{
		StartLine: startLine,
		StartCol:  startCol,
		EndLine:   endLine,
		EndCol:    endCol,
	})

}

func splitLineAndCol(s string) (line int, col int) {
	var err error

	lineCol := strings.Split(s, ".")

	line, err = strconv.Atoi(lineCol[0])

	if err != nil {
		panic(err)
	}

	col, err = strconv.Atoi(lineCol[1])

	if err != nil {
		panic(err)
	}

	return line, col

}

func addContent(c CoverStruct, file string, line int, col int, count int) {

	if _, ok := c[file].Report.CountThroughtLines[line]; !ok {
		c[file].Report.CountThroughtLines[line] = make(map[int]int, 0)
	}

	c[file].Report.CountThroughtLines[line][col] += count
}
