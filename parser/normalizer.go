package parser

import (
	"git.300brand.com/coverage/parser/decoder"
)

type Normalizer interface {
	Normalize(doc decoder.Decoder) (err error)
}
