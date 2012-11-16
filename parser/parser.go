package parser

import (
	"errors"
)

func Normalize(data []byte, n Normalizer) (err error) {
	d, err := Parse(data)
	if err != nil {
		return
	}
	return n.Normalize(d)
}

func Parse(data []byte) (d Decoder, err error) {
	t, err := Type(data)
	if err != nil {
		return
	}
	return ParseType(data, t)
}

func ParseType(data []byte, t string) (d Decoder, err error) {
	if _, ok := decoders[t]; !ok {
		errors.New("Unknown decoder type: " + t)
		return
	}
	d = decoders[t].New()
	err = d.Decode(data)
	return
}

// Tests the feed against all registered decoders to determine the appropriate
// decoder to use.
func Type(data []byte) (t string, err error) {
	for t, d := range decoders {
		if err = d.Decode(data); err == nil {
			return t, nil
		}
	}
	return "", errors.New("No valid Decoder found")
}
