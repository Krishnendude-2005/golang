package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var id = 0

type task struct {
	id   int
	desc string
}

var tasks []task

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Home Page")
}
func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is About Page")
}
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	for _, task := range tasks {
		fmt.Fprintln(w, task.id, " ", task.desc)
		//w.Write([]byte(task.desc))
	}
}
func FindTaskHandler(w http.ResponseWriter, r *http.Request) {
	reqidstr := r.PathValue("id")
	reqid, _ := strconv.Atoi(reqidstr)

	for _, task := range tasks {
		if task.id == reqid {
			fmt.Fprintln(w, task.id, " ", task.desc, "FOUND")
		}
	}
}
func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	description := r.PathValue("description")
	id += 1
	newtask := task{
		id:   id,
		desc: description,
	}
	tasks = append(tasks, newtask)
	fmt.Fprintln(w, newtask.id, " ", newtask.desc, "ADDED")
}

func main() {

	http.HandleFunc("GET /", HomePageHandler)
	http.HandleFunc("GET /task", TaskHandler)
	http.HandleFunc("GET /task/{id}", FindTaskHandler)
	http.HandleFunc("POST /task/add/{description}", PostTaskHandler)

	//always same
	fmt.Println("Server is running on PORT : 8080")
	http.ListenAndServe(":8080", nil)
}
