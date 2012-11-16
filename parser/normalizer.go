package parser

import (
	"git.300brand.com/coverage/parser"
)

type Normalizer interface {
	Normalize(doc parser.Decoder) (err error)
}
