package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	//get the file through the terminal / output
	logFile := os.Args[1] //example go run main.go - we will get main.go

	//open the file
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println("There is some problem opening the file")
	}
	//we will use differ to close the file when we are done with everything --
	defer file.Close()

	//define the counter variables for error, warning and info
	var errCount, warningCount, infoCount, total int

	//now we want to read the file
	// include scanner--using bufio reding line by line---use for loop while scanning is continued
	scanner := bufio.NewScanner(file)

	//string array called lines to store each line of that file and using range to iterate over it
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		//line := scanner.Text() //read the current lines
		total++ // random count variable to count all no of lines

		//checking if the file is a part of error/warning/info --using HasPrefix of strings inbuilt method
		switch {
		case strings.HasPrefix(line, "ERROR"):
			errCount++
		case strings.HasPrefix(line, "WARNING"):
			warningCount++
		case strings.HasPrefix(line, "INFO"):
			infoCount++
		}

	}

	out := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("There are %d errors\n", errCount)
	fmt.Printf("There are %d warnings\n", warningCount)
	fmt.Printf("There are %d infos\n", infoCount)
	fmt.Println("Analyzed at : ", out)

}
