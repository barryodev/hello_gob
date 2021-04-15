package hello_gob

import (
	"bytes"
	"encoding/gob"
	"errors"
	"log"
)

type Point struct {
	X, Y int
}

func (p Point) PointToBytes() []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	err := enc.Encode(p)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	return network.Bytes()
}

func BytesToPoint(input []byte) (Point, error) {
	var network bytes.Buffer
	network.Write(input)
	dec := gob.NewDecoder(&network)

	var output Point
	err := dec.Decode(&output)
	if err != nil {
		return Point{}, errors.New("Failed to decode bytes: " + err.Error())
	}
	return output, nil

}


