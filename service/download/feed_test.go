package download

import (
	"testing"
	"time"
)

func TestDownloadFeed(t *testing.T) {
	if _, err := downloadFeed("http://www.theregister.co.uk/headlines.atom", time.Minute); err != nil {
		t.Error(err)
	}
}

func TestDownloadFail(t *testing.T) {
	if _, err := downloadFeed("http://www.theregister.co.uk/headlines.atom", time.Millisecond); err == nil {
		t.Error("Timeout expected")
	}
}
