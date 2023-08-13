package fakename

//go:generate go run ./cmd/go-france generate

func DefaultSet(opts ...NamesetOption) (fn, ln Nameset) {
	return staticNames(opts...)
}
