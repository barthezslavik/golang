package main

import (
	"fmt"
	"github.com/barthezslavik/golang/dive/lib"
)

type Object struct {
	Content struct {
		Name string
	}
}

func main() {
	object := &Object{}
	object.Content.Name = "Text"
	fmt.Println(object)

	lib.Add("dr", "Dart")
	fmt.Println(lib.Get("dr"))
	languages := lib.GetAll()
	for _, v := range languages {
		fmt.Println(v)
	}
}
