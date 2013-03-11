package logger

import (
	"fmt"
	"time"
)

type Entries []Entry

type Entry struct {
	Time    time.Time
	Level   Level
	Message string
}

func (es *Entries) Add(e *Entry) {
	*es = append(*es, *e)
}

func (es *Entries) Debug(v ...interface{}) {
	es.WriteEntry(&Entry{
		Time:  time.Now(),
		Level: Ldebug,
	}, v...)
}

func (es *Entries) Error(err error) error {
	es.WriteEntry(&Entry{
		Time:  time.Now(),
		Level: Lerror,
	}, err)
	return err
}

func (es *Entries) WriteEntry(e *Entry, v ...interface{}) {
	switch val := v[0].(type) {
	case string:
		e.Message = fmt.Sprintf(val, v[1:]...)
	default:
		e.Message = fmt.Sprintf("%+v", v)
	}
	es.Add(e)
}
