package fakename

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type Loader interface {
	Read() (Nameset, error)
}

func NewLoader(r io.Reader) Loader { return loader{r: r} }

type loader struct {
	r io.Reader
}

func (f loader) Read() (Nameset, error) {
	records, err := csv.NewReader(f.r).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv: read all: %w", err)
	}
	dataset := &nameset{
		name:  make([]string, len(records)),
		count: make([]int, len(records)),
	}
	for i, row := range records {
		count, err := strconv.Atoi(row[1])
		if err != nil {
			return nil, fmt.Errorf("line %d: col 2: strconv: atoi: %w", i+1, err)
		}
		dataset.name[i], dataset.count[i] = row[0], count
		dataset.cum += count
	}
	return dataset, nil
}
