package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Find all .fpy files in the current directory
	files, err := filepath.Glob("*.fpy")
	if err != nil {
		fmt.Printf("Error finding .fpy files: %v\n", err)
		os.Exit(1)
	}
	if len(files) == 0 {
		fmt.Println("No .fpy files found in the current directory.")
		os.Exit(0)
	}

	// Ask the user which file to transpile
	fmt.Println("Which file would you like to transpile?")
	for i, file := range files {
		fmt.Printf("%d. %s\n", i+1, file)
	}
	var inputIndex int
	_, err = fmt.Scanln(&inputIndex)
	if err != nil {
		fmt.Printf("Error reading user input: %v\n", err)
		os.Exit(1)
	}
	if inputIndex < 1 || inputIndex > len(files) {
		fmt.Printf("Invalid selection: %d\n", inputIndex)
		os.Exit(1)
	}
	inputFile := files[inputIndex-1]

	// Read the input file into memory
	inputBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Convert "fun" to "def"
	inputCode := string(inputBytes)
	outputCode := strings.Replace(inputCode, "fun", "def", -1)

	// Write the output to a new file with ".py" extension
	outputFile := strings.TrimSuffix(inputFile, ".fpy") + ".py"
	err = ioutil.WriteFile(outputFile, []byte(outputCode), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Transpiled code written to %v\n", outputFile)
}
