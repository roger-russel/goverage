package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fullFileName := "../../tests/fixture/case-one.cov"
	ReadFile(fullFileName)

}

func ReadFile(fullFileName string) {
	file, err := os.Open(fullFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		checkLine([]string{
			"mode: ",
			"github.com/roger-russel/example-fasthttp/",
		}, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkLine(prefix []string, line string) {
	correct := false

	for _, p := range prefix {
		correct = strings.HasPrefix(line, p)
		if correct {
			break
		}
	}

	if correct {
		fmt.Println(line)
	}
}
