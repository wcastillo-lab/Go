package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"crypto/sha256"
	"bytes"
	"path/filepath"
)

func foncforimage(filepath string) []byte{
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	h := sha256.New()
	h.write(content)
	return h.Sum(nil)
	
	
}

func main() {

	image_1 := foncforimage("image_1.jpg")
	image_2 := foncforimage("image_2.jpg")
	image_3 := foncforimage("image_3.jpg")
	
	comparaison_1 := bytes.Compare(image_1,image_2)
	comparaison_2 := bytes.Compare(image_2,image_3)
	comparaison_3 := bytes.Compare(image_3,image_1)

	if comparaison_1 != 0 && comparaison_3 != 0 {
		fmt.Println("image 1 est unique")
	}
	else { 
		fmt.Println("l'image 1  n'est pas unique")
	}
	if comparaison_1 != 0 && comparaison_2 != 0 {
                fmt.Println("image 2 est  unique")
        }
	else {
                fmt.Println("l'image 2  n'est pas unique")
	}

	if comparaison_2 != 0 && comparaison_3 != 0 {
                fmt.Println("image 3 est  unique")
        }
	else {
                fmt.Println("l'image 3  n'est pas unique")
	}
}
