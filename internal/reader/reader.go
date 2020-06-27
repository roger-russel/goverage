package reader

import (
	"bufio"
	"log"
	"os"
)

func read(fileName string) (*bufio.Scanner, *os.File) {

	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile

}

func ReadTemplate() {

}
