package parser

import (
	"errors"
	"git.300brand.com/coverage/feed"
	"git.300brand.com/coverage/feed/parser/decoder"
)

func Normalize(data []byte, n feed.Normalizer) (err error) {
	d, err := Parse(data)
	if err != nil {
		return
	}
	return n.Normalize(d)
}

func Parse(data []byte) (d decoder.Decoder, err error) {
	t, err := Type(data)
	if err != nil {
		return
	}
	return ParseType(data, t)
}

func ParseType(data []byte, t string) (d decoder.Decoder, err error) {
	if _, ok := decoder.Decoders[t]; !ok {
		errors.New("Unknown decoder type: " + t)
		return
	}
	d = decoder.Decoders[t].New()
	err = d.Decode(data)
	return
}

// Tests the feed against all registered decoders to determine the appropriate
// decoder to use.
func Type(data []byte) (t string, err error) {
	for t, d := range decoder.Decoders {
		if err = d.Decode(data); err == nil {
			return t, nil
		}
	}
	return "", errors.New("No valid Decoder found")
}
