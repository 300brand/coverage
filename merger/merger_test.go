package merger

import (
	"testing"
)

type T struct {
	A         int
	B         string
	C         []byte
	D         []int
	Changelog Changelog
}

func TestHasChangelog(t *testing.T) {
	a := struct{ A int }{}
	if _, err := getChangelog(a.A); err == nil {
		t.Error("a-type does not have Changelog")
	}
	b := T{}
	if _, err := getChangelog(b); err != nil {
		t.Error("T has Changelog")
	}
	c := struct{ Changelog }{}
	if _, err := getChangelog(c); err != nil {
		t.Error("c-type has embedded Changelog")
	}
}

func TestMergeIntoEmpty(t *testing.T) {
	a := &T{}
	b := &T{}
	b.A = 5
	b.Changelog.Changed("A")

	changed, err := Merge(a, b)
	if err != nil {
		t.Error(err)
	}
	if !changed {
		t.Error("No changes made")
	}
	if a.A != 5 {
		t.Error("Change not applied")
		t.Errorf("%+v", a)
	}
}

func TestTypeMatch(t *testing.T) {
	a := struct{ A int }{}
	b := struct{ A int }{}
	c := struct{ B int }{}
	if !typesMatch(a, b) {
		t.Error("Types of a and b should match")
	}
	if typesMatch(b, c) {
		t.Error("Types of b and c should not match")
	}
	if typesMatch(a, c) {
		t.Error("Types of a and c should not match")
	}
}
