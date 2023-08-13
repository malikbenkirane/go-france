package fakename

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	for _, test := range []struct {
		name string
		src  string
	}{
		{
			name: "lastname",
			src:  "data/noms_sorted_by_count.csv",
		},
		{
			name: "firstname",
			src:  "data/prenoms_sorted_by_count.csv",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(test.src)
			require.NoError(t, err)
			ns, err := NewSet(f)
			require.NoError(t, err)
			for i := 0; i < 100; i++ {
				t.Log(ns.RandomName())
			}
		})
	}
}
