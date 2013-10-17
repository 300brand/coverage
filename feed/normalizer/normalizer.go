package normalizer

import (
	"github.com/300brand/coverage/feed/parser/decoder"
)

type Normalizer interface {
	Normalize(doc decoder.Decoder) (err error)
}
