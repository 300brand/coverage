package parser

type Decoder interface {
	New() Decoder
	Decode([]byte) error
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
