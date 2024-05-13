package digdug

import (
	"os"

	"go.uber.org/dig"
)

// Builder provides chainable variants to the [Container] setup methods.
type Builder struct {
	Container
}

// DefaultBuilder provides a global Builder using a [dig.Container]
var DefaultBuilder = New()

// Provide teaches the [Container] how to build values of one or more types and
// expresses their dependencies. If an error is returned, the Builder will panic.
func (b *Builder) Provide(constructor any, opts ...dig.ProvideOption) *Builder {
	if err := b.Container.Provide(constructor, opts...); err != nil {
		b.quit(err)
	}

	return b
}

// Decorate provides a decorator for a type that has already been provided in
// the [Container]. If an error is returned, the Builder will panic.
func (b *Builder) Decorate(constructor any, opts ...dig.DecorateOption) *Builder {
	if err := b.Container.Decorate(constructor, opts...); err != nil {
		b.quit(err)
	}

	return b
}

// Scope creates a new Builder with a new [dig.Scope], using the given name and
// options from current [Container].
func (b *Builder) Scope(name string, opts ...dig.ScopeOption) *Builder {
    scope := b.Container.Scope(name, opts...)
    return &Builder{scope}
}

func (b *Builder) quit(err error) {
    if c, ok := b.Container.(*dig.Container); ok {
	    _ = dig.Visualize(c, os.Stderr, dig.VisualizeError(err))
    }

	panic(err)
}

// New creates a Builder with a new [dig.Container]
func New() *Builder {
	return &Builder{dig.New()}
}
