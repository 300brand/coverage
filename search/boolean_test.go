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

func TestBooleanBadMatches(t *testing.T) {
	tests := []struct {
		Haystack string
		Needle   string
	}{
		{
			"a b c",
			"d",
		},
		{
			"a b c",
			"a AND d",
		},
		{
			"a b c",
			"d OR e",
		},
		{
			"a b c",
			"e OR d OR f",
		},
		{
			"a b c",
			"d OR a AND e",
		},
		{
			"a b c",
			"e OR d OR f AND h AND g",
		},
		{
			"a b c d",
			"d c AND b a",
		},
		// Addition of NOT
		{
			"a b c",
			"a NOT b",
		},
		{
			"a b c",
			"a OR b NOT c",
		},
		{
			"a b c",
			"a NOT c OR d OR e",
		},
		// NOT tests from Jamie
		{
			"emcee",
			"EMC NOT emcee OR emceeing",
		},
		{
			"collision damage waiver",
			"CDW NOT collision damage waiver",
		},
	}
	for i, test := range tests {
		b := NewBoolean(test.Needle)
		if b.Match([]byte(test.Haystack)) {
			t.Errorf("[%d] '%s' should not be found in '%s'", i, test.Needle, test.Haystack)
		}
	}
}

func TestBooleanGoodMatches(t *testing.T) {
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
		{
			"a b c d",
			"a b AND b c",
		},
		{
			"a b c d",
			"a b AND c d",
		},
		{
			"a b c",
			"a AND b OR c AND d",
		},
		// Addition of NOT
		{
			"a b c",
			"a NOT d",
		},
	}
	for i, test := range tests {
		b := NewBoolean(test.Needle)
		if !b.Match([]byte(test.Haystack)) {
			t.Errorf("[%d] '%s' not found in '%s'", i, test.Needle, test.Haystack)
		}
	}
}

func TestBooleanMinTerms(t *testing.T) {
	tests := []struct {
		Query string
		Terms int
	}{
		{"EMC OR Brocade", 1},
		{"PaaS AND EMC OR Brocade", 1},
		{"PaaS AND EMC OR PaaS AND Brocade", 2},
		{"Platform as a Service AND EMC OR Platform as a Service AND Brocade", 3},
	}
	for i, test := range tests {
		b := NewBoolean(test.Query)
		if n := b.MinTerms(); n != test.Terms {
			t.Errorf("[%d] %s: %d != %d", i, test.Query, n, test.Terms)
		}
	}
}

func TestBooleanTerms(t *testing.T) {
	tests := []struct {
		Query string
		Terms []string
	}{
		{"spacemonkey NOT NASA", []string{"spacemonkey"}},
	}
	for i, test := range tests {
		b := NewBoolean(test.Query)
		t.Logf("[%d] %v", i, b.Terms())
	}
}
