package search

import (
	"testing"
)

func TestBooleanTree(t *testing.T) {
	tests := []struct {
		Search string
		Tree   [][]string
	}{
		{
			"phrase",
			[][]string{[]string{"phrase"}},
		},
		{
			"some phrase OR another phrase",
			[][]string{
				[]string{"some phrase"},
				[]string{"another phrase"},
			},
		},
		{
			"some phrase AND another phrase",
			[][]string{[]string{"some phrase", "another phrase"}},
		},
		{
			"a AND b OR c",
			[][]string{[]string{"a", "b"}, []string{"c"}},
		},
	}
	for n, test := range tests {
		b := NewBoolean(test.Search)
		if len(test.Tree) != len(b.Tree) {
			t.Errorf("[%d] Tree length mismatch", n)
			t.Logf("test %+v", test.Tree)
			t.Logf("bool %+v", b.Tree)
			continue
		}
		for i := range test.Tree {
			if len(test.Tree[i]) != len(b.Tree[i]) {
				t.Errorf("[%d][%d] Subtree length mismatch", n, i)
				t.Logf("test %+v", test.Tree)
				t.Logf("bool %+v", b.Tree)
				continue
			}
			for j := range test.Tree[i] {
				if x, y := test.Tree[i][j], b.Tree[i][j].String(); x != y {
					t.Errorf("[%d][%d] Inequal values '%s' != '%s'", i, j, x, y)
				}
			}
		}
	}
}

func TestBooleanMatches(t *testing.T) {
	tests := []struct {
		Haystack string
		Needle   string
	}{
		{
			"a b c",
			"a",
		},
		{
			"a b c",
			"a AND b",
		},
		{
			"a b c",
			"a AND c",
		},
		{
			"a b c",
			"a OR b",
		},
		{
			"a b c",
			"d OR b",
		},
		{
			"a b c",
			"e OR d OR c",
		},
		{
			"a b c",
			"d OR a AND c",
		},
		{
			"a b c",
			"e OR d OR b AND c AND a",
		},
		{
			"a b c",
			"a AND c OR d",
		},
	}
	for i, test := range tests {
		b := NewBoolean(test.Needle)
		if !b.Match([]byte(test.Haystack)) {
			t.Errorf("[%d] '%s' not found in '%s'", i, test.Needle, test.Haystack)
		}
	}
}
