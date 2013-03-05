package lexer

import (
	"testing"
)

type tests map[string]string

var benchmarkString = []byte("This string, while short, tests Normalize's abilities!")

func BenchmarkNormalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Normalize(benchmarkString)
	}
}

func TestApostrophes(t *testing.T) {
	runTests(t, tests{
		"test's": "test's",
	})
}

func TestNewLines(t *testing.T) {
	runTests(t, tests{
		"test":         "test",
		"\ntest":       " test",
		"test\n":       "test ",
		"test\r":       "test ",
		"test\r\n":     "test  ",
		"test\ntest":   "test test",
		"test\r\ntest": "test  test",
	})
}

func TestPunctuation(t *testing.T) {
	runTests(t, tests{
		"test? test":  "test test",
		"test, test":  "test test",
		"test! test":  "test test",
		"test. test":  "test test",
		"test & test": "test  test",
	})
}

func TestQuotes(t *testing.T) {
	runTests(t, tests{
		`"test`:       "test",
		`"test"`:      "test",
		`"test test"`: "test test",
	})
}

func TestSpaces(t *testing.T) {
	runTests(t, tests{
		"test":             "test",
		"test test":        "test test",
		"test test test":   "test test test",
		" test test test":  " test test test",
		"test test test ":  "test test test ",
		"test  test  test": "test  test  test",
	})
}

func TestTabs(t *testing.T) {
	runTests(t, tests{
		"test":       "test",
		"\ttest":     " test",
		"test\t":     "test ",
		"test\ttest": "test test",
	})
}

func runTests(t *testing.T, ts tests) {
	for s, expect := range ts {
		n := Normalize([]byte(s))
		if string(n) != expect {
			t.Errorf("Expected %s; Got %s for `%s'", expect, n, s)
		}
	}
}
