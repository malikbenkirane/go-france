package fakename

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Nameset interface {
	RandomName() string
	Label() string
	All() (name []string, count []int, cum int)
}

type nsconf struct {
	r     *rand.Rand
	label string
}

func defaultNamesetConfig() nsconf {
	return nsconf{
		r:     rand.New(rand.NewSource(time.Now().UnixNano())),
		label: "no label",
	}
}

type NamesetOption func(nsconf) nsconf

func NamesetWithRand(r *rand.Rand) NamesetOption {
	return func(n nsconf) nsconf {
		n.r = r
		return n
	}
}

func NamesetWithLabel(label string) NamesetOption {
	return func(n nsconf) nsconf {
		n.label = label
		return n
	}
}

func NewSet(r io.Reader, opts ...NamesetOption) (Nameset, error) {
	l := NewLoader(r)
	_ns, err := l.Read()
	ns := _ns.(*nameset)
	if err != nil {
		return nil, fmt.Errorf("ns: read: %w", err)
	}
	ns.conf = defaultNamesetConfig()
	for _, opt := range opts {
		ns.conf = opt(ns.conf)
	}
	return ns, nil
}

type nameset struct {
	conf nsconf

	name  []string
	count []int
	cum   int
}

func (ns nameset) RandomName() string {
	p := ns.conf.r.Float64()
	pos := int(p * float64(ns.cum))
	var i, count, cur int
	for i, count = range ns.count {
		cur += count
		if cur >= pos &&
			(ns.name[i] != "AUTRES NOMS" && ns.name[i] != "_PRENOMS_RARES") {
			break
		}
	}
	return ns.name[i]
}

func (ns nameset) Label() string {
	return ns.conf.label
}

func (ns nameset) All() (name []string, count []int, cum int) {
	return ns.name, ns.count, ns.cum
}
