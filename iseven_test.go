package iseven_test

import (
	"strconv"
	"testing"

	"github.com/aoisensi/go-iseven"
)

var cases = []int{0, 1, 2, 5, 686, 47545, 54612, 153818, 999999}

func FuzzIsEven(f *testing.F) {
	for _, n := range cases {
		f.Add(n)
	}
	f.Fuzz(func(t *testing.T, n int) {
		if 0 > n || n >= 1000000 {
			t.Skip()
		}
		client := iseven.NewClient(nil)
		resp, err := client.IsEven(n)
		if err != nil {
			t.Fatal(err)
		}
		if resp.IsEven != ((n % 2) == 0) {
			t.Fail()
		}
	})
}

func TestIsEven(t *testing.T) {
	for _, n := range cases {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			client := iseven.NewClient(nil)
			resp, err := client.IsEven(n)
			if err != nil {
				t.Fatal(err)
			}
			if resp.IsEven != ((n % 2) == 0) {
				t.Fail()
			}
		})
	}
}
