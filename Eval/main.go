package main


import (
	"fmt"
	"reflect"
)

func main() {
	
	task := []string{}


	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)
	http.ListenAndServe(":8080", nil)
}



type Task struct {
        Description string
        Done        bool
}

func modifSlice(sl []string, idx int){
	
}

