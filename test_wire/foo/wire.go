//go:build wireinject

package foo

import (
	"github.com/google/wire"
)

func provideMyFoo() *MyFoo {
	b := new(MyFoo)
	*b = "Hello, World!"
	return b
}

var Set = wire.NewSet(
	wire.Bind(new(Foo), new(*MyFoo)),
	provideMyFoo,
	NewBar,
)

func InitializeBar() *MyBar {
	wire.Build(Set)
	return nil
}
