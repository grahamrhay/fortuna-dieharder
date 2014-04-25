Outputs an infinite sequence of random bytes to stdout:

    go run src/fortuna-dieharder.go

Useful for running the dieharder test suite against this [implementation](https://github.com/seehuhn/fortuna) of the Fortuna PRNG:

    go run src/fortuna-dieharder.go | dieharder -a -g 200
