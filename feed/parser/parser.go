package parser

import (
	"fmt"
	"git.300brand.com/coverage/feed/parser/decoder"
	"strings"
)

func Parse(data []byte) (d decoder.Decoder, err error) {
	t, err := Type(data)
	if err != nil {
		return
	}
	return ParseType(data, t)
}

func ParseType(data []byte, t string) (d decoder.Decoder, err error) {
	if _, ok := decoder.Decoders[t]; !ok {
		err = fmt.Errorf("Unknown decoder type: %s", t)
		return
	}
	d = decoder.Decoders[t].New()
	err = d.Decode(data)
	return
}

// Tests the feed against all registered decoders to determine the appropriate
// decoder to use.
func Type(data []byte) (t string, err error) {
	errs := make([]string, 0, len(decoder.Decoders))
	for t, d := range decoder.Decoders {
		if err = d.Decode(data); err != nil {
			errs = append(errs, fmt.Sprintf("[%s]: %s", t, err))
			continue
		}
		return t, nil
	}
	return "", fmt.Errorf("parser.Type: No decoder found. (%s)", strings.Join(errs, "; "))
}
