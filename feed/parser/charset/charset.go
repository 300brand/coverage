package charset

import (
	"bytes"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"encoding/xml"
	"fmt"
	"github.com/300brand/coverage/feed/parser/decoder"
)

func TryAll(doc decoder.Decoder, data []byte) (err error) {
	dec := xml.NewDecoder(bytes.NewReader(data))
	dec.CharsetReader = charset.NewReader

	// Use default charset / charset from document
	if err = dec.Decode(doc); err == nil {
		return
	}

	// Try all charsets in arsenal
	errs := make([]error, len(charset.Names()))
	for i, name := range charset.Names() {
		r, err := charset.NewReader(name, bytes.NewReader(data))
		if err != nil {
			return err
		}
		dec := xml.NewDecoder(r)
		if err = dec.Decode(doc); err == nil {
			return nil
		}
		errs[i] = fmt.Errorf("[%s] %s", name, err)
	}
	return fmt.Errorf("Could not find a good charset: %v", errs)
}
