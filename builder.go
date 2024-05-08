package digdug

import (
	"os"

	"go.uber.org/dig"
)

// Builder provides chainable variants to the [dig.Container] setup methods.
type Builder struct {
	*dig.Container
}

// Provide teaches the [dig.Container] how to build values of one or more types and
// expresses their dependencies. If an error is returned, the Builder will panic.
func (b *Builder) Provide(constructor any, opts ...dig.ProvideOption) *Builder {
	if err := b.Container.Provide(constructor, opts...); err != nil {
		b.quit(err)
	}

	return b
}

// Decorate provides a decorator for a type that has already been provided in
// the [dig.Container]. If an error is returned, the Builder will panic.
func (b *Builder) Decorate(constructor any, opts ...dig.DecorateOption) *Builder {
	if err := b.Container.Decorate(constructor, opts...); err != nil {
		b.quit(err)
	}

	return b
}

func (b *Builder) quit(err error) {
	_ = dig.Visualize(b.Container, os.Stderr, dig.VisualizeError(err))
    panic(err)
}

// New creates a Builder with a new [dig.Container]
func New() *Builder {
	return &Builder{dig.New()}
}
