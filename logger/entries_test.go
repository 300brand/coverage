package logger

import (
	"bytes"
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	e := Entries{}
	e.Debug(errors.New("Test Error"))
	if msg := e[0].Message; msg != "Test Error" {
		t.Error("Invalid message:", msg)
	}
}

func TestErrorPassthru(t *testing.T) {
	e := Entries{}
	in := errors.New("Test Error")
	out := e.Error(in)
	if out == nil {
		t.Error("No error returned")
	}
	if out != in {
		t.Error("Invalid error returned")
	}

}

func TestIOLogging(t *testing.T) {
	w := &bytes.Buffer{}
	EnableLogging(w)
	e := Entries{}
	e.Debug("Test")
	e.Error(errors.New("Test Error"))
	if w.String() == "" {
		t.Error("Empty buffer returned")
	}
	DisableLogging()
}

func TestLength1(t *testing.T) {
	e := Entries{}
	e.Debug("Test")
	if len(e) != 1 {
		t.Error("Invalid entry count:", len(e))
	}
}

func TestLength2(t *testing.T) {
	e := Entries{}
	e.Debug("Test1")
	e.Debug("Test2")
	if len(e) != 2 {
		t.Error("Invalid entry count:", len(e))
	}
}

func TestString(t *testing.T) {
	e := Entries{}
	e.Debug("Test")
	if msg := e[0].Message; msg != "Test" {
		t.Error("Invalid message:", msg)
	}
}

func TestStringf(t *testing.T) {
	e := Entries{}
	e.Debug("Testing %d %d %d", 1, 2, 3)
	if msg := e[0].Message; msg != "Testing 1 2 3" {
		t.Error("Invalid message:", msg)
	}
}

func TestStruct(t *testing.T) {
	e := Entries{}
	e.Debug(struct{ A int }{A: 1})
	if msg := e[0].Message; msg != "[{A:1}]" {
		t.Error("Invalid message:", msg)
	}
}
