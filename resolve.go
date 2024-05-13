package digdug

// Resolve resolves the parameterized type from the [Container], returning an
// error if not found.
func Resolve[T any](c Container) (T, error) {
	ch := make(chan T, 1)
	defer close(ch)

	if err := c.Invoke(func(value T) {
		ch <- value
	}); err != nil {
		return *new(T), err
	}

	return <-ch, nil
}

// MustResolve resolves the parameterized type from the [Container], panicking
// if not found.
func MustResolve[T any](c Container) T {
	value, err := Resolve[T](c)
	if err != nil {
		panic(err)
	}

	return value
}
