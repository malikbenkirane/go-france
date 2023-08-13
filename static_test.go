package fakename

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultSet(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	fn, ln := DefaultSet(NamesetWithRand(r))
	assert.Equal(t, "MARCELLE", fn.RandomName(), "random firstname")
	assert.Equal(t, "MARTIN", ln.RandomName(), "random lastname")
}
