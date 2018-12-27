package main

import (
	"fmt"

	"github.com/ATTHDEV/jlib"
)

func main() {

	json := jlib.NewObject(jlib.Object{
		"page": 1,
		"fruits": jlib.Array{
			"apple", "peach",
		},
		"drink": jlib.Object{
			"name": "coffee",
		},
	})

	fmt.Println("show pretty json string...")
	fmt.Println(json.ToPrettyString())

	json.Put("drink", json.Object("drink").Put("name", "cocoa"))
	fmt.Println("show json data after change coffe to cocoa..")
	fmt.Println(json.ToPrettyString())

}
