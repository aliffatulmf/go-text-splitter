package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	wordsPerChunk := flag.Int("words", 2000, "maximum number of words per chunk")
	inputFile := flag.String("input", "", "path to input markdown file")
	outputDir := flag.String("output-dir", "output", "path to output directory")
	flag.Parse()

	// Check input file type is valid
	fileExt, err := CheckFileType(*inputFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Open the input file
	input, err := OpenFile(*inputFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer CloseFile(input)

	// Create output directory if it doesn't exist
	err = os.MkdirAll(*outputDir, 0755)
	if err != nil {
		fmt.Printf("Error creating output directory: %s\n", errors.Unwrap(err))
		os.Exit(1)
	}

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(input)

	// Create a file to write the current chunk to
	var outputFile *os.File
	// Keep track of the number of words in the current chunk
	var wordCount int
	// Keep track of the number of output files written
	var fileCount int

	// Read the input file line by line
	for scanner.Scan() {
		// Split the line into words
		lineWords := strings.Split(scanner.Text(), " ")
		for _, word := range lineWords {
			// If the maximum number of words is reached, close the current file and start a new file with the next chunk
			if wordCount == *wordsPerChunk {

				CloseFile(outputFile)
				fileCount++
				wordCount = 0
			}

			// If a new file needs to be created, create the file and write the current chunk to it
			if wordCount == 0 {
				outputFilePath := filepath.Join(*outputDir, fmt.Sprintf("output-%d%s", fileCount, fileExt))
				outputFile, err = os.Create(outputFilePath)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
			}

			// Write the current word to the current chunk
			outputFile.WriteString(word + " ")
			wordCount++
		}
		// Write a line break to the current chunk to preserve the original formatting
		outputFile.WriteString("\n")
	}
	CloseFile(outputFile)

	for i := 0; i <= fileCount; i++ {
		outputFilePath := filepath.Join(*outputDir, fmt.Sprintf("output%d%s", i, fileExt))
		fmt.Println(outputFilePath)
	}
}

func CheckFileType(file string) (string, error) {
	fileExt := filepath.Ext(file)
	switch fileExt {
	case ".md", ".txt":
		return fileExt, nil
	default:
		return fileExt, fmt.Errorf("input file must be a .md or .txt file")
	}
}

func OpenFile(filePath string) (*os.File, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening input file: %s\n", errors.Unwrap(err))
	}
	if fileInfo.IsDir() {
		return nil, fmt.Errorf("Error opening input file: %s is a directory\n", filePath)
	}
	input, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Error opening input file: %s\n", errors.Unwrap(err))
	}
	return input, nil
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Printf("error closing file: %w", err)
		os.Exit(1)
	}
}
