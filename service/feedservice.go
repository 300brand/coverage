package service

import (
	"github.com/300brand/coverage"
)

type FeedService interface {
	Update(*coverage.Feed) error
}
