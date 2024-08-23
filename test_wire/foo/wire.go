//go:build wireinject

package foo

import (
	"github.com/google/wire"
)

func InitializeBar(foo Foo) *MyBar {
	wire.Build(NewBar)
	return nil
}
