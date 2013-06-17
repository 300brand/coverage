package main

import (
	"git.300brand.com/coverage/skytypes"
	"net/http"
)

type Manager struct{}

func init() {
	s.RegisterService(new(Manager), "")
}

func (m *Manager) AddOneFeed(r *http.Request, in *skytypes.NullType, out *skytypes.NullType) (err error) {
	cmd := &skytypes.ClockCommand{
		Command: "once",
	}
	return GetService("Manager").Send(nil, "QueueFeedAdder", cmd, skytypes.Null)
}

func (m *Manager) ProcessNextFeed(r *http.Request, in *skytypes.ClockCommand, out *skytypes.ClockResult) (err error) {
	return
}
