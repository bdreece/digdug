package digdug

import "go.uber.org/dig"

// Resolve resolves the parameterized type from the [dig.Container], returning an
// error if not found.
func Resolve[T any](c *dig.Container) (T, error) {
    ch := make(chan T, 1)
    defer close(ch)
    
    if err := c.Invoke(func(value T) {
        ch<- value
    }); err != nil {
        return *new(T), err
    }

    return <-ch, nil
}
