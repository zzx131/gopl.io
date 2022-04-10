package main

import (
	"bufio"
	"fmt"
	"os"
)

// 文件写入
func main() {
	outPutFile, outputErr := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputErr != nil {
		fmt.Printf("an error occurred with file opening or creation\n")
		return
	}
	defer outPutFile.Close()
	outputWriter := bufio.NewWriter(outPutFile)
	outputString := "hello word \n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
