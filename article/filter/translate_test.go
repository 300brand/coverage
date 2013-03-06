package filter

import (
	"testing"
)

func TestTranslation(t *testing.T) {
	s := "WASHINGTON DC—And the party… rocked! “We’ve partied so hard,” said one attendee"
	expect := "WASHINGTON DC-And the party... rocked! \"We've partied so hard,\" said one attendee"
	translated := translateString(s)
	if translated != expect {
		t.Logf("Expect: %s", expect)
		t.Logf("Got:    %s", translated)
		t.Error("Translation failure")
	}
}
