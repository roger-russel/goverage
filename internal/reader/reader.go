package reader

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

//CoverStruct       filename    line   column count
type CoverStruct map[string]map[int]map[int]int

//ReadFile read an atomic coverage file and turn it into a map
func ReadFile(filename string) CoverStruct {

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	mode, _ := readLine(reader)

	if mode != "mode: atomic" {
		panic("only model atomic is supported")
	}

	coverStruct := make(CoverStruct, 0)

	for rawLine, next := readLine(reader); next; {
		line := strings.Split(rawLine, ":")
		file := line[0]
		splitContent(coverStruct, file, line[1])
	}

	return coverStruct

}

func readLine(reader *bufio.Reader) (string, bool) {

	str, _, err := reader.ReadLine()

	if err == io.EOF {
		return "", false
	}

	return strings.TrimRight(string(str), "\r\n"), true

}

func splitContent(c CoverStruct, file string, lineRaw string) {

	swapRaw := strings.Split(lineRaw, " ")
	count, err := strconv.Atoi(swapRaw[2])

	if err != nil {
		panic(err)
	}

	swapRawLineCol := strings.Split(swapRaw[0], ",")

	startLine, startCol := splitLineAndCol(swapRawLineCol[0])
	endLine, endCol := splitLineAndCol(swapRawLineCol[1])

	addContent(c, file, startLine, startCol, count)
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

	if _, ok := c[file]; !ok {
		c[file] = make(map[int]map[int]int, 0)
	}

	if _, ok := c[file][line]; !ok {
		c[file][line] = make(map[int]int, 0)
	}

	c[file][line][col] += count

}
