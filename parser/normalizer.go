package parser

type Normalizer interface {
	Normalize(doc Decoder) (err error)
}
