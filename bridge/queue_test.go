package bridge

import (
	"git.300brand.com/coverage/config"
	"io/ioutil"
	"net/http"
	"testing"
)

func testGet(t *testing.T) {
	r, err := http.Get("http://" + config.RPC.Address)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	t.Logf("%s", body)
}

func TestQueue(t *testing.T) {
	items, err := Queue()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("items: %+v", items)
}
