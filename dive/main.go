package main

import (
	"fmt"
	"github.com/barthezslavik/golang/dive/domain"
)

func main() {
	order := &domain.Customer{}
	order.Id = 1
	order.Name = "Order"
	fmt.Println(order)
}
