package main

import (
	"fmt"
	"test_wire/foo"
)

func provideMyFoo() *foo.MyFoo {
	b := new(foo.MyFoo)
	*b = "Hello, World!"
	return b
}

func main() {
	myFoo := provideMyFoo()
	b := foo.InitializeBar(myFoo)
	fmt.Println(b.Run())
}
