package lexer

import (
	"strings"
	"testing"
)

type tests map[string]string

func TestApostrophes(t *testing.T) {
	runTests(t, tests{
		"test's": "test's",
	})
}

func TestNewLines(t *testing.T) {
	runTests(t, tests{
		"test":         "test",
		"\ntest":       "test",
		"test\n":       "test",
		"test\r":       "test",
		"test\r\n":     "test",
		"test\ntest":   "testtest",
		"test\r\ntest": "testtest",
	})
}

func TestQuotes(t *testing.T) {
	runTests(t, tests{
		`"test`:       "test",
		`"test"`:      "test",
		`"test test"`: "testtest",
	})
}

func TestSpaces(t *testing.T) {
	runTests(t, tests{
		"test":             "test",
		"test test":        "testtest",
		"test test test":   "testtesttest",
		" test test test":  "testtesttest",
		"test test test ":  "testtesttest",
		"test  test  test": "testtesttest",
	})
}

func TestTabs(t *testing.T) {
	runTests(t, tests{
		"test":       "test",
		"\ttest":     "test",
		"test\t":     "test",
		"test\ttest": "testtest",
	})
}

func runTests(t *testing.T, ts tests) {
	for s, cleaned := range ts {
		w := GetWords([]byte(s))
		joined := strings.Join(w.Strings(), "")
		if joined != cleaned {
			t.Errorf("Expected %s; Got %s for `%s'", cleaned, joined, s)
		}
	}
}
