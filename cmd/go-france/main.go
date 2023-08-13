package main

import (
	"fmt"
	"os"

	"github.com/malikbenkirane/go-france/cmd/go-france/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
