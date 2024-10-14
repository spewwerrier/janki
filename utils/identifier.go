package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
)

// generates a random identifier. Concept taken from
// maybe I should use this with some modification to make API keys
// http://antirez.com/news/99
func GenerateIdentifier(counter int64) string {
	// generate a random seed
	seedRand := rand.Reader
	seed, err := rand.Int(seedRand, big.NewInt(1000000000000000000))
	if err != nil {
		log.Panic(err)
	}

	s := sha1.New()

	// convert seed to bytes
	data := new(bytes.Buffer)
	binary.Write(data, binary.LittleEndian, seed.Int64())

	// convert counter to bytes
	data2 := new(bytes.Buffer)
	err = binary.Write(data2, binary.LittleEndian, counter)
	if err != nil {
		log.Panic(err)
	}

	// append those bytes
	param := data.Bytes()
	param = append(param, data2.Bytes()...)

	s.Write(param)

	res := s.Sum(nil)
	v := fmt.Sprintf("%x", res)
	return v
}

func main() {
	v := GenerateIdentifier(16)
	fmt.Println(v)
}
