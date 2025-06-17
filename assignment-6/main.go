package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Logger interface {
	Log(message string)
}
type ConsoleLogger struct {
	//can take user input
}
type FileLogger struct {
	//can take user input
}
type RemoteLogger struct {
	//can take user input
}

// Log method of console
func (c ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message, "!")
}

// Log method of file
func (f FileLogger) Log(message string) {
	fmt.Println("File:", message, "!")
}

// Log method of Remote
func (r RemoteLogger) Log(message string) {
	fmt.Println("Remote:", message, "!")
}

func logAll(logitems []Logger, message string) {
	for _, value := range logitems {
		value.Log(strings.TrimSpace(message))
	}
}
func main() {
	fmt.Println("Interface Type 2 : Interface with Composition and Polymorphism")

	var consolesms string
	var filesms string
	var remotesms string
	fmt.Println("Enter message for Console")
	fmt.Scan(&consolesms)

	fmt.Println("Enter message for File")
	fmt.Scan(&filesms)

	fmt.Println("Enter message for Remote")
	fmt.Scan(&remotesms)

	//creating instance of Logger
	var l Logger

	//craeting an slice of Logger interface to show same sms for all
	var logitems []Logger

	//creating instance for Console
	c := ConsoleLogger{}
	l = c
	l.Log(consolesms)
	logitems = append(logitems, l) //appending values of l for each type

	f := FileLogger{}
	l = f
	l.Log(filesms)
	logitems = append(logitems, l) ////appending values of l for each type

	r := RemoteLogger{}
	l = r
	l.Log(remotesms)
	logitems = append(logitems, l) ////appending values of l for each type

	//calling the logAll function - passing the Logger slice , iterate over it and call each time with message passed

	fmt.Println("Enter the message you want to show as same for everyone")
	reader := bufio.NewReader(os.Stdin) //creating reader for input from user
	inputsms, _ := reader.ReadString('\n')
	logAll(logitems, inputsms)

}
