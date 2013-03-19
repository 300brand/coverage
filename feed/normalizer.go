package feed

import (
	"git.300brand.com/coverage/feed/parser/decoder"
)

type Normalizer interface {
	Normalize(doc decoder.Decoder) (err error)
}
