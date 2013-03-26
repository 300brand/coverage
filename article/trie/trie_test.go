package trie

import (
	"testing"
)

func TestAddSingle(t *testing.T) {
	tr := New()
	tr.Add("test")
	if !tr.Trunk.End {
		t.Error("Expected trunk End to be true")
	}
}

func TestAddBigSmall(t *testing.T) {
	tr := New()
	tr.Add("testing")
	tr.Add("tests")
	if !tr.Trunk.Branches[int('i'-CharLow)].End {
		t.Error("Expected 'i' End to be true")
	}
	if !tr.Trunk.Branches[int('s'-CharLow)].End {
		t.Error("Expected 's' End to be true")
	}
}

func TestAddSmallBig(t *testing.T) {
	tr := New()
	tr.Add("tests")
	tr.Add("testing")
	if !tr.Trunk.Branches[int('i'-CharLow)].End {
		t.Error("Expected 'i' End to be true")
	}
	if !tr.Trunk.Branches[int('s'-CharLow)].End {
		t.Error("Expected 's' End to be true")
	}
}

func TestAddTestFirst(t *testing.T) {
	tr := New()
	tr.Add("test")
	tr.Add("testing")
	tr.Add("tests")
	if !tr.Trunk.End {
		t.Error("Expected trunk End to be true")
	}
	if !tr.Trunk.Branches[int('i'-CharLow)].End {
		t.Error("Expected 'i' End to be true")
	}
	if !tr.Trunk.Branches[int('s'-CharLow)].End {
		t.Error("Expected 's' End to be true")
	}
}

func TestAddTestLast(t *testing.T) {
	tr := New()
	tr.Add("testing")
	tr.Add("tests")
	tr.Add("test")
	if !tr.Trunk.End {
		t.Error("Expected trunk End to be true")
	}
	if !tr.Trunk.Branches[int('i'-CharLow)].End {
		t.Error("Expected 'i' End to be true")
	}
	if !tr.Trunk.Branches[int('s'-CharLow)].End {
		t.Error("Expected 's' End to be true")
	}
}

func TestIndex(t *testing.T) {
	//                 Core:                     Extended:
	indexes := []rune("abcdefghijklmnopqrstuvwxyzA'._B09482")
	b := NewBranch()
	for i, r := range indexes {
		if idx := b.MakeIndex(r); idx != i {
			t.Errorf("Invalid index for %s: %d, expected %d", string([]rune{r}), idx, i)
		}
	}
}

func TestAccommodate(t *testing.T) {
	b := NewBranch()
	if l := len(b.Branches); l != 0 {
		t.Errorf("Expected zero-len Branches. Got: %d", l)
	}
	b.Accommodate(b.MakeIndex('a'))
	if l := len(b.Branches); l != 1 {
		t.Errorf("Expected 1 Branch. Got: %d", l)
	}
	b.Accommodate(b.MakeIndex('j'))
	if l := len(b.Branches); l != 10 {
		t.Errorf("Expected 10 Branches. Got: %d", l)
	}
	b.Accommodate(b.MakeIndex('0'))
	if l := len(b.Branches); l != 27 {
		t.Errorf("Expected 27 Branches. Got: %d", l)
	}
}

func TestDump(t *testing.T) {
	tr := New()
	tr.Add("tease")
	tr.Add("teases")
	tr.Add("teased")
	tr.Add("teaser")
	tr.Add("tests")
	tr.Add("tested")
	tr.Add("test")
	tr.Add("testing")
	t.Logf("\n%s", tr.Dump())
}
