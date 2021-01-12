package main

import (
	"fmt"
	"flag"
)

func main() {
	const VERSION = "1.0"
	theversion := flag.Bool("version",false, "show the program version")
	flag.Parse()
	if *theversion {
		fmt.Println(VERSION)
		return
	}
	fmt.Println("Hello, World")
	
}
  



