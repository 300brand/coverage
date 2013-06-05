package main

import (
	"errors"
	"git.300brand.com/coverage/skytypes"
	"github.com/skynetservices/skynet"
	"time"
)

func (s *Clock) Article(ri *skynet.RequestInfo, in *skytypes.ClockCommand, out *skytypes.ClockResult) (err error) {
	return
}

func (s *Clock) Feed(ri *skynet.RequestInfo, in *skytypes.ClockCommand, out *skytypes.ClockResult) (err error) {
	return
}

func (s *Clock) QueueFeedAdder(ri *skynet.RequestInfo, in *skytypes.ClockCommand, out *skytypes.NullType) (err error) {
	return s.processCommand(s.Tickers["QueueFeedAdder"], in)
}

func (s *Clock) queueFeedAdder() {
	id := &skytypes.ObjectId{}
	if err := c.GetService("Queue", "", "", "").Send(nil, "AddFeed", skytypes.Null, id); err != nil {
		s.Log.Error(err.Error())
		return
	}
	s.Log.Trace("Added to feed queue: " + id.Id.Hex())
}

func (s *Clock) processCommand(t *Ticker, cmd *skytypes.ClockCommand) (err error) {
	switch cmd.Command {
	case "once":
		t.Once <- true
	default:
		err = errors.New("Unknown command: " + cmd.Command)
	}
	return
}

func runner(t *Ticker) {
	for {
		select {
		case <-t.Once:
			t.F()
		case <-t.Start:
			t.Ticker = time.NewTicker(t.Tick)
		case <-t.Stop:
			t.Ticker.Stop()
		case <-t.Ticker.C:
			t.F()
		}
	}
}
