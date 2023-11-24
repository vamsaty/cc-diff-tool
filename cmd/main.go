package main

import (
	"fmt"
	ccUtils "github.com/vamsaty/cc-utils"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatalln("Usage: cc-diff-tool file1 file2")
	}

	// Check if the files exist
	for _, fileName := range args[1:3] {
		if !ccUtils.FileExists(fileName) {
			log.Fatalln("File does not exist: ", fileName)
		}
	}

	file1, err := ccUtils.ReadFile(args[1])
	if err != nil {
		log.Fatalln("Error reading file: ", args[1])
	}
	file2, err := ccUtils.ReadFile(args[2])
	if err != nil {
		log.Fatalln("Error reading file: ", args[2])
	}

	output := ExecuteDiff(
		strings.Split(string(file1), "\n"),
		strings.Split(string(file2), "\n"),
	)
	fmt.Println(output)
}
