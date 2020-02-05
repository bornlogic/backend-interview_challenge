# IBCC

Ibcc - Interview backend code challenge - is a reposity used for **bornlogic** challenge code

## Matrix

Matrix is a lib for perform some tests over matrix.

### Tests

- Is Triangular
- Is Lower Triangular
- Is Upper Triangular
- Is Diagonal
- Is Square

## testMatrix

testMatrix is a cli tool for test over matrix passed as stdin or by arg.

format used is csv, example:
```
[[1,2,3],        1,2,3
 [1,2,3],   -->  1,2,3
 [1,2,3]]        1,2,3
```

### Usage

`exit 1` -> means test fails

`exit 0` -> means test success or no test passed

for more you can use `--help` flag:
```
Usage of testMatrix:
  -m value
        matrix in csv format (shorthand)
  -matrix value
        matrix in csv format
  -t    test if given matrix is triangular (shorthand)
  -triangular
        test if given matrix is triangular
  -v    prints feedback of test in stderr (shorthand)
  -verbose
        prints feedback of test in stderr
```

### Examples

Stdin - matrix is triangular
```sh
$ echo -e '0,5\n0,0' | go run cmd/testMatrix/main.go --triangular
```
exit status 0


Stdin - matrix is not triangular
```sh
$ echo -e '0,5\n1,0' | go run cmd/testMatrix/main.go --triangular
```
exit status 1


Arg - matrix is not triangular (verbose)
```bash
$ go run cmd/testMatrix/main.go -t --matrix $'0,5\n1,0' -v
testMatrix: 2020/02/04 16:42:59 matrix is not triangular
```
exit status 1


Arg - no test passed (verbose)
```bash
$ go run cmd/testMatrix/main.go -m $'0,5\n0,0' -v
testMatrix: 2020/02/04 23:45:31 no test passed
```
exit status 0


## Api

You can use tests of matrix by api.

For this you send a json for `/testMatrix` api endpoint with test name and matrix for the server test.

### Schema

schema used in body for test matrix

spec: https://github.com/NeowayLabs/jsonschema/blob/master/docs/spec.md

```json
{
	"matrix": {
		"type": "array",
		"required": true,
		"format": {
			"type": "array",
			"format": {
				"type": "int",
			}
		}
	},
	"testName": {
		"type": "string",
		"required": true,
	}
}
```

### Feedback

any case where is not a success returns a quoted string with description about error

- 200 - Valid matrix
- 400 - Bad request (missing or typo)
- 406 - Not Acceptable (invalid matrix, test fails)

### Examples

Valid triangular matrix - returns 200 without message
```sh
$ curl http://127.0.0.1:3000/testMatrix --data '{"matrix": [[1,2,3], [0,1,1], [0,0,1]], "testName": "triangular"}'
```

Invalid triangular matrix - returns 406 and a quoted message
```sh
$ curl http://127.0.0.1:3000/testMatrix --data '{"matrix": [[1,2,3], [0,1,1], [1,0,1]], "testName": "triangular"}' -v
"matrix test failed: is not triangular"
```

Invalid test name - returns 406 and a quoted message
```sh
$ curl http://127.0.0.1:3000/testMatrix --data '{"matrix": [[1,2,3], [0,1,1], [0,0,1]], "testName": "invalid"}'
"matrix test failed: invalid test name"
```

Invalid body - returns 400 and a quoted message
```sh
$ curl http://127.0.0.1:3000/testMatrix --data '{' -v
"invalid body: unexpected end of JSON input"
```

Missing matrix - returns 400 and a quoted message
```sh
$ curl http://127.0.0.1:3000/testMatrix --data '{"matrix":[], "testName": "triangular"}' -v
"missing matrix"
```

Missing test name - returns 400 and a quoted message
```sh
$ curl http://127.0.0.1:3000/testMatrix --data '{"matrix": [[1,2,3], [0,1,1], [1,0,1]], "testName": ""}' -v
"missing test name"
```
