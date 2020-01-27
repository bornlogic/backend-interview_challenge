package main

import (
	"os"
	"log"
	"flag"
	"io/ioutil"

	"github.com/iuryfukuda/ibcc/matrix"
)

var Matrix matrix.Matrix
var TestIsTriangular bool
var IsVerbose bool

func main(){
	flag.Parse()
	log.SetPrefix("testMatrix: ")

	if Matrix == nil {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("can't read stdin: %s", err)
		}
		Matrix.Set(string(b))
	}

	if len(Matrix) == 0 {
		log.Fatal("missing matrix")
	}

	switch {
	case TestIsTriangular:
		if !matrix.IsTriangular(Matrix) {
			os.Exit(1)
		} else if IsVerbose {
			log.Print("matrix is triangular")
		}
	}
}

func init() {
	const usageMatrix = "matrix in csv format"
	flag.Var(&Matrix, "matrix", usageMatrix)
	flag.Var(&Matrix, "m", usageMatrix+" (shorthand)")

	const (
		usageTestIsTriangular = "test if given matrix is triangular"
		defaultTestIsTriangular = false
	)
	flag.BoolVar(
		&TestIsTriangular, "triangular",
		defaultTestIsTriangular, usageTestIsTriangular,
	)
	flag.BoolVar(
		&TestIsTriangular, "t",
		defaultTestIsTriangular, usageTestIsTriangular+" (shorthand)",
	)

	const (
		usageIsVerbose = "prints feedback of test in stderr"
		defaultIsVerbose = false
	)
	flag.BoolVar(&IsVerbose, "verbose", defaultIsVerbose, usageIsVerbose)
	flag.BoolVar(&IsVerbose, "v", defaultIsVerbose, usageIsVerbose+" (shorthand)")
}


