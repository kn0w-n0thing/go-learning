package foo

type MyBar struct {
	foo Foo
}

func NewBar(foo Foo) *MyBar {
	return &MyBar{foo: foo}
}

func (bar *MyBar) Run() string {
	return bar.foo.Foo()
}
