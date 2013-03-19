package decoder

type Decoder interface {
	New() Decoder
	Decode([]byte) error
}

var Decoders = map[string]Decoder{}

// Adds a Decoder to the decoder map. You cannot add a decoder type of the same
// name if it already exists.
func RegisterDecoder(t string, d Decoder) {
	if _, ok := Decoders[t]; ok {
		panic("Cannot re-add decoder: " + t)
	}
	Decoders[t] = d
}
