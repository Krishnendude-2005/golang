package main

import "fmt"

type Logger interface {
	Log()
}

type File struct {
	s string
}

func (f *File) Log() {
	if f == nil {
		fmt.Println("Nil pointer but still works")
		return
	}
	fmt.Println("Logging from value receiver but pointer receives")
}

func (f *File) Save() {
	if f == nil {
		fmt.Println("Nil pointer but still works")
	}
	fmt.Println("Saving from pointer receiver")
}

func main() {
	var l Logger
	var file = File{"Hello World"}
	l = &file
	l.Log()

	//f1 := File{}
	//f2 := &File{}

	//l = f1
	//l.Log()
	//
	//l = f2
	//l.Log()
}
