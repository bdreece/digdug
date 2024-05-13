// Package digdug provides wrapper utilities around [dig.Container].
package digdug

import "go.uber.org/dig"

// Container provides an interface definition that corresponds with
// the setup methods of [dig.Container] and [dig.Scope].
type Container interface {
    Decorate(decorator interface{}, opts ...dig.DecorateOption) error
    Invoke(function interface{}, opts ...dig.InvokeOption) error
    Provide(constructor interface{}, opts ...dig.ProvideOption) error
    Scope(name string, opts ...dig.ScopeOption) *dig.Scope
}

