package digdug

// Supply creates a constructor function for the provided value. For example:
//
//	type Config struct { /* omitted for brevity */ }
//
//	func Foo(cfg *Config) {
//	    container := digdug.New().
//	        Provide(digdug.Supply(cfg)).
//	        Provide(/* ... */).
//	        Container
//	}
func Supply[T any](value T) func() T {
	return func() T {
		return value
	}
}
