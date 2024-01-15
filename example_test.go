package vegagoja_test

import (
	"context"
	"fmt"

	"github.com/xo/vegagoja"
)

func Example() {
	buf, err := vegagoja.Render(context.Background(), spec, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len: %d", len(buf))
	// Output:
	// len: 0
}

var (
	spec string
	data string
)
