package main

import (
	"fmt"

	"github.com/ATTHDEV/jlib"
)

func main() {

	json := jlib.NewObject(`{"page": 1, "fruits": ["apple", "peach"] , "drink" : {"name" : "coffee"}}`)

	fmt.Println(json.Int("page"))

	fruits := json.Array("fruits")
	for _, f := range fruits.ToArray() {
		fmt.Println(f)
	}

	for k, v := range json.Object("drink").ToMap() {
		fmt.Println(k, v)
	}

}
