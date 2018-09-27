package main

import (
	"fmt"
	"strings"
)
import "log"

const k = 1

var v float64 = 9.9

type S struct {
	name string
}

func f(n int) {
	fmt.Printf("f says %d\n", n)
}
func (s S) m() {
	fmt.Printf("Hi %s\n", s.name)
}
