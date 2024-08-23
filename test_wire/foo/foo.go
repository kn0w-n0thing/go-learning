package foo

type Foo interface {
	Foo() string
}

type MyFoo string

func (b *MyFoo) Foo() string {
	return string(*b)
}
