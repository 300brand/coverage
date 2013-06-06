package main

import (
	"git.300brand.com/coverage/skytypes"
	"net/http"
)

type Manager struct{}

func init() {
	s.RegisterService(new(Feed), "")
}

func (m *Manager) ProcessNextFeed(r *http.Request, in *skytypes.ClockCommand, out *skytypes.ClockResult) (err error) {
	return
}
