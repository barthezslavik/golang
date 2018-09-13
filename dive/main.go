package main

import "fmt"

type Object struct {
	Content struct {
		Name string
	}
}

func main() {
	object := &Object{}
	object.Content.Name = "Text"
	fmt.Println(object)
}
