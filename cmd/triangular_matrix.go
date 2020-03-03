package cmd

import (
	"flag"
	"os"

	"github.com/marciusvinicius/Interview-Backend-Code-Challenge/pkg/matrix"
)

func WhereIsTriangular(rows [][]int) string {
	builder := matrix.Builder{}
	for r := range rows {
		builder = append(make(matrix.Builder, r))
	}

	m1, err := matrix.Build(builder)

	if err != nil {
		return "You passed a 0x0 or 1x0 matrix."
	}

	return m1.TypeOfMatrix()
}

func main() {
	textPtr := flag.String("text", "", "Text to parse. (Required)")
	flag.Parse()
	if *textPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	print(textPtr)
}
