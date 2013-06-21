package main

import (
	"errors"
	"fmt"
	"git.300brand.com/coverage"
	"git.300brand.com/coverage/skytypes"
	"github.com/skynetservices/skynet"
	"time"
)

func (s *Manager) Article(ri *skynet.RequestInfo, in *skytypes.ClockCommand, out *skytypes.ClockResult) (err error) {
	return
}

func (s *Manager) FeedProcessor(ri *skynet.RequestInfo, in *skytypes.ClockCommand, out *skytypes.ClockResult) (err error) {
	return s.processCommand(s.Tickers["FeedProcessor"], in)
}

func (s *Manager) feedProcessor() (err error) {
	in := &coverage.Feed{}
	out := &coverage.Feed{}

	if err = s.Queue.Send(nil, "NextFeed", skytypes.Null, out); err != nil {
		s.Log.Error(err.Error())
		return
	}
	s.Log.Trace(fmt.Sprintf("%s Got ID", out.ID.Hex()))

	*in = *out
	if err = s.FeedDownload.Send(nil, "Download", in, out); err != nil {
		s.Log.Error(out.Log.Error(err).Error())
		s.StorageWriter.Send(nil, "SaveFeed", out, in)
		return
	}
	s.Log.Trace(fmt.Sprintf("%s Downloaded", out.ID.Hex()))

	*in = *out
	if err = s.FeedProcess.Send(nil, "Process", in, out); err != nil {
		s.Log.Error(out.Log.Error(err).Error())
		s.StorageWriter.Send(nil, "SaveFeed", out, in)
		return
	}
	s.Log.Trace(fmt.Sprintf("%s Processed", out.ID.Hex()))

	*in = *out
	if err = s.StorageWriter.Send(nil, "SaveFeed", in, out); err != nil {
		s.Log.Error(err.Error())
		return
	}
	s.Log.Trace(fmt.Sprintf("%s Saved", out.ID.Hex()))
	return
}

func (s *Manager) QueueFeedAdder(ri *skynet.RequestInfo, in *skytypes.ClockCommand, out *skytypes.NullType) (err error) {
	return s.processCommand(s.Tickers["QueueFeedAdder"], in)
}

func (s *Manager) queueFeedAdder() (err error) {
	id := &skytypes.ObjectId{}
	if err = s.Queue.Send(nil, "AddFeed", skytypes.Null, id); err != nil {
		s.Log.Error(err.Error())
		return
	}
	s.Log.Trace("Added to feed queue: " + id.Id.Hex())
	return
}

func (s *Manager) processCommand(t *Ticker, cmd *skytypes.ClockCommand) (err error) {
	switch cmd.Command {
	case "once":
		t.Once <- true
		return t.F()
	case "start":
		t.Start <- true
	case "stop":
		t.Stop <- true
	default:
		err = errors.New("Unknown command: " + cmd.Command)
	}
	return
}

func runner(t *Ticker) {
	for {
		select {
		case <-t.Once:
			//t.F()
		case <-t.Start:
			t.Ticker = time.NewTicker(t.Tick)
		case <-t.Stop:
			t.Ticker.Stop()
		case <-t.Ticker.C:
			t.F()
		}
	}
}
