package parser

import (
	"errors"
	"net/url"
	"time"
)

type Feed struct {
	Title    string
	Articles []Article
}

type Article struct {
	Published time.Time
	Title     string
	URL       url.URL
}

type Decoder interface {
	Decode([]byte) error
	Feed() Feed
}

var decoders = map[string]Decoder{}

// Adds a Decoder to the decoder map. You cannot add a decoder type of the same
// name if it already exists.
func RegisterDecoder(t string, d Decoder) {
	if _, ok := decoders[t]; ok {
		panic("Cannot re-add decoder: " + t)
	}
	decoders[t] = d
}

// Parses the incoming bytes to spit out the appropriate Feed. The second
// argument may have the appropriate Decoder type to use, or a blank string to
// automatically determine which decoder to use.
func Parse(data []byte, t string) (f Feed, err error) {

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
