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

func (m *Manager) ProcessNextFeed(r *http.Request, in *skytypes.NullType, out *skytypes.NullType) (err error) {
	svc := GetService("Manager")
	svc.SetTimeout(0, 5*time.Minute)
	cmd := &skytypes.ClockCommand{
		Command: "once",
	}
	return svc.Send(nil, "FeedProcessor", cmd, skytypes.Null)
}

func (m *Manager) StartQueue(r *http.Request, in *skytypes.NullType, out *skytypes.ClockCommand) (err error) {
	out = &skytypes.ClockCommand{
		Command: "start",
		Tick:    time.Second * 10,
	}
	return GetService("Manager").Send(nil, "QueueFeedAdder", out, skytypes.Null)
}

func (m *Manager) StopQueue(r *http.Request, in *skytypes.NullType, out *skytypes.NullType) (err error) {
	cmd := &skytypes.ClockCommand{
		Command: "stop",
	}
	return GetService("Manager").Send(nil, "QueueFeedAdder", cmd, skytypes.Null)
}

func (m *Manager) StartFeeds(r *http.Request, in *skytypes.NullType, out *skytypes.ClockCommand) (err error) {
	out = &skytypes.ClockCommand{
		Command: "start",
		Tick:    time.Second * 10,
	}
	return GetService("Manager").Send(nil, "FeedProcessor", out, skytypes.Null)
}

func (m *Manager) StopFeeds(r *http.Request, in *skytypes.NullType, out *skytypes.NullType) (err error) {
	cmd := &skytypes.ClockCommand{
		Command: "stop",
	}
	return GetService("Manager").Send(nil, "FeedProcessor", cmd, skytypes.Null)
}
