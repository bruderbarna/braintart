package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func getSource(path string) ([]byte, error) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("Couldn't open source file for reading. Aborting.")
	}
	return input, nil
}

func generateOutput(source []byte) []byte {
	output := []byte("#include<stdio.h>\n#include<stdlib.h>\nint main() { char *p = calloc(30000, 1); char *psave = p; if (!p) return 1; ")
	for i := 0; i < len(source); i++ {
		switch source[i] {
		case 'i':
			output = incrementPointer(output)
		case 'd':
			output = decrementPointer(output)
		case 'I':
			output = incrementValueAtPointer(output)
		case 'D':
			output = decrementValueAtPointer(output)
		case 'p':
			output = printValueAtPointer(output)
		case 'g':
			output = getAndStoreInputAtPointer(output)
		case 'f':
			output = jumpForwardIfValueAtPointerIsZero(output)
		case 'b':
			output = jumpBackwards(output)
		}
	}
	output = append(output, []byte("free(psave); return 0; }")...)
	return output
}

func writeOutput(filepath string, output []byte) error {
	err := ioutil.WriteFile(filepath, output, 0755)
	if err != nil {
		return errors.New("Couldn't open output file for writing. Aborting")
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Number of arguments isn't 2. First argument should be the source file, second should be the output file")
	}

	inputFilepath := os.Args[1]
	outputFilepath := os.Args[2]
	input, err := getSource(inputFilepath)
	if err != nil {
		log.Fatal(err)
	}

	output := generateOutput(input)

	if err := writeOutput(outputFilepath, output); err != nil {
		log.Fatal(err)
	}
}
