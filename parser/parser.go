package parser

import (
	"errors"
	"net/url"
	"time"
)

// Parses the incoming bytes to spit out the appropriate Feed. The second
// argument may have the appropriate Decoder type to use, or a blank string to
// automatically determine which decoder to use.
func Parse(data []byte, t string) (f Feed, err error) {
	if t == "" {
		if t, err = Type(data); err != nil {
			return
		}
	}
	doc := decoders[t].New()
	if err = doc.Decode(data); err != nil {
		return
	}
	f = doc.Feed()
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
