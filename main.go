//package main
//
//import (
//	"bufio"
//	"fmt"
//	"math"
//	"os"
//	"strings"
//)
//
//// Currency Map
//var rates = map[string]float64{
//	"USD": 1.0,
//	"INR": 83.12,
//	"EUR": 13.0,
//	"JPY": 10.0,
//}
//
//// Variables Needed
//var amount float64
//var source_currency string
//var target_currency string
//
//// Function for Taking Input
//func takeInput() {
//	fmt.Println("Enter Amount")
//	fmt.Scan(&amount)
//
//	fmt.Println("Enter Source Currency")
//	fmt.Scan(&source_currency)
//
//	fmt.Println("Enter Target Currency")
//	fmt.Scan(&target_currency)
//
//}
//
//// Function for Main Conversion Logic
//func convert() float64 {
//	return (amount / rates[source_currency]) * rates[target_currency]
//}
//
//// Main Function
//func main() {
//	var s float64
//	fmt.Println(s)
//
//	sliceExample()
//	structexample()
//	//switch {
//	//case 1 == 1:
//	//	return 1
//	//
//	//case 2 == 2:
//	//	return 2
//	//
//	//}
//
//	fmt.Println("Hello World")
//
//	practice()
//
//	power := pow(3, 2, 10)
//	fmt.Println(power)
//	//----------------------------------------Currency Convertor-----------------------------
//	takeInput()
//
//	var totalSum float64
//	totalSum = convert()
//
//	fmt.Println(amount, source_currency, "is equivalent to", totalSum, target_currency)
//}
//
//// ----------------------------------------------------------------------------------------------
//func practice() {
//	//Multiple If Statements
//	var variable1 = 10
//	var variable2 = 20
//	if variable1 == 10 && variable2 == 20 {
//		fmt.Println("Vaiables are equals to", variable1, "and", variable2)
//	}
//}
//
//func pow(x, y, lim float64) float64 {
//	if v := math.Pow(x, y); v < lim {
//		fmt.Printf("value of v is %g :: within lim %g\n", v, lim)
//		return v
//	} else {
//		fmt.Printf("value of v is %g :: outside lim %g\n", v, lim)
//		return lim
//	}
//}
//
//// -----------------Struct example-----------------------------------------------------
//func structexample() {
//	type employee struct {
//		name     string
//		age      int
//		salary   int
//		position string
//	}
//	type Vertex struct {
//		X, Y int
//	}
//
//	///making a struct
//	//type -1
//	var employee1 employee
//	employee1.name = "Krish"
//	employee1.age = 20
//	employee1.salary = 30000
//	employee1.position = "SDE-1"
//
//	//type -2
//	employee2 := employee{"krishnendu", 20, 30000, "SDE-1"}
//
//	fmt.Println(employee1)
//	fmt.Println(employee2)
//
//	p := &employee2
//	p.salary = 70000
//	fmt.Println(employee2)
//
//	//struct literals
//	var (
//		v1      = Vertex{1, 2}  // has type Vertex
//		v2      = Vertex{X: 1}  // Y:0 is implicit
//		v3      = Vertex{}      // X:0 and Y:0
//		pointer = &Vertex{1, 2} // has type *Vertex
//	)
//	fmt.Println(v1, v2, v3, pointer)
//}
//
//// -----------------------------------------------------------------------------------
//// -----------------Slice example-----------------------------------------------------
//func sliceExample() {
//	arr := []int{1, 2, 3, 4, 5}
//	fmt.Println(arr)
//
//	//creating/taking slice from this array
//	slicearr := arr[0:3]
//	fmt.Println(slicearr)
//
//	//appending in present sliced array
//	slicearr = append(slicearr, 20, 30)
//	fmt.Println(slicearr)
//
//	fmt.Println(len(slicearr))
//	fmt.Println(cap(slicearr))
//
//	//Other slices that share the same underlying array will see those changes.
//}

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
