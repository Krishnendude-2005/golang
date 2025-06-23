package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	fibonacci := fibo()
	//fibonacci calling
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci())
	}

	dummyfunc := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Printf("dummy func: %f\n", dummyfunc(5, 12))

	fmt.Println("function passed as parameter : ", functionpassing(dummyfunc))

	mapexample()
	//struct array example----------------------------
	s := []struct {
		name string
		age  int
	}{
		{"krish", 20},
		{"Hello", 10},
	}

	fmt.Println(s)
	//-------------------------------------------------------

	//slice exercise
	sliceExercise(10, 20)

}

// Slice Exercise -- https://go.dev/tour/moretypes/18-------------------------------
func sliceExercise(dx, dy int) {
	dummySlice := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		rowslice := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			rowslice[x] = uint8(x * y)
		}
		dummySlice[y] = rowslice
	}

	for i, j := range dummySlice {
		fmt.Println("index", i, " ", "values", j)
	}
} //------------------------------------------------------------

// map--
func mapexample() {
	m := make(map[int]string)

	m[1] = "hello"
	m[2] = "world"

	fmt.Println(m)

	//map word count-----------------------------
	wordcount := func(str string) map[string]int {
		m := make(map[string]int)
		words := strings.Fields(str)

		for _, word := range words {
			m[word]++
		}
		return m
	}

	fmt.Println(wordcount("Hello Namaste namaste Hello"))
	//--------------------------------------------------------
}

func functionpassing(fn func(float64, float64) float64) float64 {
	return fn(5, 12)
}

// fibonacci closure
func fibo() func() int {
	num1, num2 := 0, 1

	return func() int {
		result := num1
		num1, num2 = num2, num1+num2
		return result
	}
}
