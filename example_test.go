package iseven_test

import (
	"fmt"

	"github.com/aoisensi/go-iseven"
)

func Example() {
	// Create iseven client
	client := iseven.NewClient(nil)

	// Check to see if this number is even
	resp, err := client.IsEven(686)
	if err != nil {
		panic(err)
	}

	// Print out result
	fmt.Println(resp.IsEven) // true
}
