package service

import (
	"git.300brand.com/coverage"
)

type FeedService interface {
	Update(*coverage.Feed) error
}
