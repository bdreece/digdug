package digdug

// Supply creates a constructor function for the provided value. For example:
//
//	type Config struct { /* omitted for brevity */ }
//
//	func main() {
//      var cfg Config
//      
//      /* parse config from file */
//
//	    container := digdug.New().
//	        Provide(digdug.Supply(&cfg)).
//	        Provide(/* ... */).
//	        Container
//	}
func Supply[T any](value T) func() T {
	return func() T {
		return value
	}
}
