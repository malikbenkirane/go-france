package fakename

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestExampleNewSet(t *testing.T) {
	ExampleNewSet()
}

func ExampleNewSet() {

	// Loading both firstnames and lastnames concurrently.

	fn, ln, err := func() (Nameset, Nameset, error) {
		dst := make(chan Nameset, 2)

		{
			var eg errgroup.Group
			for _, _s := range []struct {
				src   string
				label string
			}{
				{
					src:   "data/noms_sorted_by_count.csv",
					label: "lastnames",
				},
				{
					src:   "data/prenoms_sorted_by_count.csv",
					label: "firstnames",
				},
			} {
				s := _s
				eg.Go(func() error {
					f, err := os.Open(s.src)
					if err != nil {
						return fmt.Errorf("%s: os: open: %w", s.label, err)
					}
					ns, err := NewSet(f, NamesetWithRand(rand.New(rand.NewSource(42)))) // read also go doc math/rand
					if err != nil {
						return fmt.Errorf("%s: new set: %w", s.label, err)
					}
					dst <- ns
					return nil
				})
			}

			if err := eg.Wait(); err != nil {
				return nil, nil, err
			}

		}

		{

			var fn, ln Nameset

			ns := <-dst
			if ns.Label() == "lastnames" {
				ln = ns
				fn = <-dst
			} else {
				fn = ns
				ln = <-dst
			}

			return fn, ln, nil

		}

	}()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fn.RandomName(), ln.RandomName())

	// Output: MARCELLE LEMEE

}
