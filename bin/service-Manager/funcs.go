package main

import (
	"errors"
	//"fmt"
	//"git.300brand.com/coverage"
	"git.300brand.com/coverage/skytypes"
	"github.com/skynetservices/skynet"
	"time"
)

func (s *Manager) Article(ri *skynet.RequestInfo, in *skytypes.ClockCommand, out *skytypes.ClockResult) (err error) {
	return
}

// func (s *Manager) ProcessNextFeed(ri *skynet.RequestInfo, in *skytypes.ClockCommand, out *skytypes.ClockResult) (err error) {
// 	return s.processCommand(s.Tickers["ProcessNextFeed"], in)
// }

//func (s *Manager) processNextFeed() {
//	f := &coverage.Feed{}
//	if err := s.Queue.Send(nil, "NextFeed", skytypes.Null, f); err != nil {
//		s.Log.Error(err.Error())
//		return
//	}
//	s.Log.Trace(fmt.Sprintf("%s Got ID", f.ID.Hex()))
//
//	if err = s.FeedDownload.Send(ri, "Download", in, out); err != nil {
//		return
//	}
//	s.Log.Trace(fmt.Sprintf("%s Downloaded", f.ID.Hex()))
//
//	*in = *out
//	if err = s.FeedProcess.Send(ri, "Process", in, out); err != nil {
//		return
//	}
//	s.Log.Trace(fmt.Sprintf("%s Processed", f.ID.Hex()))
//
//	*in = *out
//	if err = s.StorageWriter.Send(ri, "SaveFeed", in, out); err != nil {
//		return
//	}
//	s.Log.Trace(fmt.Sprintf("%s Saved", f.ID.Hex()))
//	return
//}

func (s *Manager) QueueFeedAdder(ri *skynet.RequestInfo, in *skytypes.ClockCommand, out *skytypes.NullType) (err error) {
	return s.processCommand(s.Tickers["QueueFeedAdder"], in)
}

func (s *Manager) queueFeedAdder() {
	id := &skytypes.ObjectId{}
	if err := s.Queue.Send(nil, "AddFeed", skytypes.Null, id); err != nil {
		s.Log.Error(err.Error())
		return
	}
	s.Log.Trace("Added to feed queue: " + id.Id.Hex())
}

func (s *Manager) processCommand(t *Ticker, cmd *skytypes.ClockCommand) (err error) {
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
