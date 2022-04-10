package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n")
		return
	}
	defer inputFile.Close()

	inputRead := bufio.NewReader(inputFile)
	for {
		inputString, readerErr := inputRead.ReadString('\n')
		fmt.Printf("the input was: %s", inputString)
		if readerErr == io.EOF {
			return
		}
	}
}
