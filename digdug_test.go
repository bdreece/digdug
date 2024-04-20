package digdug

import "testing"

type (
	foo struct{ x int }
	bar struct{ y int }
	baz struct{ z int }
)

func TestFoo(t *testing.T) {
    value := 5
	c := New().
        Provide(Supply(value)).
		Provide(func(value int) foo {
            return foo{value}
        }).
        Provide(func(foo foo) bar {
            return bar{foo.x * 2}
        }).
        Provide(func(bar bar) baz {
            return baz{bar.y / 2}
        }).
        Container

    baz, err := Resolve[baz](c)
    if err != nil {
        t.Errorf("failed to resolve baz: %v\n", err)
    }

    if baz.z != value {
        t.Errorf("baz.z != value; expected %d, got %d\n", value, baz.z)
    }
}
