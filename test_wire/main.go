package main

import (
	"fmt"
	"test_wire/foo"
)

func main() {
	b := foo.InitializeBar()
	fmt.Println(b.Run())
}
