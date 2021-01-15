package main


import (
	"fmt"
	"reflect"
)

func main() {
	
	http.HandleFunc("/", list)
        http.HandleFunc("/done", done)
        http.HandleFunc("/add", add)
        http.ListenAndServe(":8080", nil)
	

		

		
}

type Task struct {
        Description string
        Done        bool
}

type List struct {
	ID   string 
	Task string

}

var task []Task {}
func list(rw http.ResponseWriter, _ *http.Request){
	list := []List{}
	task = []Task{
		{"wi", false},
		{"su"}
}	





	


