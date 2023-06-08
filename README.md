# Text Splitter
Text-Splitter is a command-line program written in Go that splits a text file into smaller chunks based on a maximum number of words per chunk, while preserving the original formatting of the input file. The program can be used to split large text files into smaller files for easier processing.

### Usage
To use Text-Splitter, run the program from the command line with the following flags:
- `-words`: the maximum number of words per chunk (default: 100)
- `-input`: the path to the input text file
- `-output-dir`: the path to the output directory for the chunked files (default: `<input_file_directory>/output`)

Example usage:
```
./text-splitter -words 4000 -input input.txt -output-dir output
```

### How it works
Text-Splitter reads the input file line by line and splits each line into words. It then keeps track of the current number of words in the current chunk and writes each chunk to a separate file when the maximum number of words is reached. A line break is added to the current chunk to preserve the original formatting of the input file.

The program uses the following algorithm to split the input file into chunks:
1. Open the input file.
2. Create the output directory if it doesn't exist.
3. Read the input file line by line.
4. For each line, split the line into words.
5. For each word, add the word to the current chunk and keep track of the current number of words in the chunk.
6. If the maximum number of words per chunk is reached, write the current chunk to a separate file and start a new chunk with the next line.
7. If a new file needs to be created, create the file and write the current chunk to it.
8. Write a line break to the current chunk to preserve the original formatting.
9. Close the last file.
10. Print the paths of the output files.

### Dependencies
Text-Splitter has no external dependencies and can be run on any system with Go installed.

### License
Text-Splitter is released under the MIT License. See the LICENSE file for details.