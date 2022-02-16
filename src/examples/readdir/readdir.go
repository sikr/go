package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err == nil {
		for _, file := range files {
			fmt.Println(file.Name())
		}
	} else {
		fmt.Println("Error: ", err)
	}
}