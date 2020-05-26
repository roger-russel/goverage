package reader

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

//LineCover is the page coverage data
type LineCover struct {
	//NumberOfStatements in line
	NumberOfStatements int
	//        line   column count
	Report map[int]map[int]int
}

//CoverStruct       filename
type CoverStruct map[string]*LineCover

//ReadFile read an atomic coverage file and turn it into a map
func ReadFile(filename string) CoverStruct {

	readFile, err := os.Open(filename)
	defer readFile.Close()

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	if fileScanner.Scan() {
		if strings.Trim(fileScanner.Text(), "") != "mode: atomic" {
			panic("only model atomic is supported")
		}
	}

	coverStruct := make(CoverStruct, 0)
	duplicatedHash := make(map[string]map[string]bool)

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), ":")
		file := line[0]
		splitContent(coverStruct, file, line[1], duplicatedHash)
	}

	return coverStruct

}

func splitContent(c CoverStruct, file string, lineRaw string, duplicatedHash map[string]map[string]bool) {

	swapRaw := strings.Split(lineRaw, " ")
	count, err := strconv.Atoi(swapRaw[2])

	if err != nil {
		panic(err)
	}

	if _, ok := c[file]; !ok {
		c[file] = &LineCover{
			NumberOfStatements: 0,
			Report:             make(map[int]map[int]int, 0),
		}
	}

	statements, err := strconv.Atoi(swapRaw[1])

	if err != nil {
		panic(err)
	}

	if _, ok := duplicatedHash[file]; !ok {
		duplicatedHash[file] = make(map[string]bool)
	}

	if isDuplicated, ok := duplicatedHash[file][swapRaw[0]]; !(isDuplicated || ok) {
		c[file].NumberOfStatements += statements
	}

	swapRawLineCol := strings.Split(swapRaw[0], ",")

	startLine, startCol := splitLineAndCol(swapRawLineCol[0])
	endLine, endCol := splitLineAndCol(swapRawLineCol[1])

	addContent(c, file, startLine, startCol, count)
	//Because it is end of it will not increate the statement on new lines
	addContent(c, file, endLine, endCol, count*-1)

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

	if _, ok := c[file].Report[line]; !ok {
		c[file].Report[line] = make(map[int]int, 0)
	}

	c[file].Report[line][col] += count

}
