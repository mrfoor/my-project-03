package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . sample.txt result.txt")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	
	if inputFile != "sample.txt" || outputFile != "result.txt" {
		fmt.Println("ERROR")
		return
	}

	lines, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	var processedLines []string
	for _, line := range lines {
		line = Binarytodecimal(line)
		line = Hextodecimal(line)
		line = ProcessToUpper(line)
		line = ProcessCapitalize(line)
		line = ProcessLowTag(line)
		line = FixArticles(line)
		line = FixPunctuation(line)
		processedLines = append(processedLines, line)
	}

	err = writeFile(outputFile, processedLines)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("File processed successfully. Results written to", outputFile)
}
// Gère l'entrée, applique les transformations et écrit le fichier de sortie.