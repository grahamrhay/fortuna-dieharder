package main

import (
	"crypto/aes"
	"encoding/binary"
	"os"

	"github.com/seehuhn/fortuna"
)

func main() {
	rng, err := fortuna.NewRNG("")
	if err != nil {
		panic("cannot initialise the RNG: " + err.Error())
	}
	defer rng.Close()
	gen := fortuna.NewGenerator(aes.NewCipher)
	for true {
		err = binary.Write(os.Stdout, binary.LittleEndian, gen.PseudoRandomData(16))
		if err != nil {
			panic("binary.Write failed:" + err.Error())
		}
	}
}
