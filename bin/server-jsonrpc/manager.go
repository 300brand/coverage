package main

import (
	"git.300brand.com/coverage/skytypes"
	"net/http"
	"time"
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

func (m *Manager) ProcessNextFeed(r *http.Request, in *skytypes.ClockCommand, out *skytypes.NullType) (err error) {
	svc := GetService("Manager")
	svc.SetTimeout(0, 5*time.Minute)
	cmd := &skytypes.ClockCommand{
		Command: "once",
	}
	return svc.Send(nil, "FeedProcessor", cmd, skytypes.Null)
}
