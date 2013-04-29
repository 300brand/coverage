package node

import (
	"github.com/skynetservices/skynet/service"
	"testing"
)

func TestTypes(t *testing.T) {
	var _ service.ServiceDelegate = &StorageReader{}
	var _ service.ServiceDelegate = &StorageWriter{}
}
