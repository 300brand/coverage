package merger

import (
	"time"
)

type Changelog []Change

func (c *Changelog) Add(field string, when time.Time) {
	ch, ok := c.GetChange(field)
	if ok {
		ch.Times.Add(when)
	} else {
		*c = append(*c, Change{
			Field: field,
			Times: Times{when},
		})
	}
}

func (c Changelog) Changed(fields ...string) {
	now := time.Now()
	for _, field := range fields {
		c.Add(field, now)
	}
}

func (c Changelog) Fields() (names []string) {
	for _, ch := range c {
		names = append(names, ch.Field)
	}
	return
}

func (c Changelog) GetChange(field string) (*Change, bool) {
	for i := range c {
		if c[i].Field == field {
			return &c[i], true
		}
	}
	return nil, false
}

func (c Changelog) LastChangeTo(field string) time.Time {
	ch, ok := c.GetChange(field)
	if !ok {
		return time.Time{}
	}
	return ch.Times.Last()
}
