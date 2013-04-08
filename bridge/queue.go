package bridge

type Queue struct {
	Feeds []struct {
		Id  uint64
		Url string
	}
}

type queueResponse struct {
	QueueId  uint64 `json:"id"`
	Class    string `json:"class"`
	ObjectId uint64 `json:"object_id"`
	Data     string `json:"data"`
}

func GetQueue(LastId, Limit int) (resp []queueResponse, err error) {
	b := New()
	defer b.Close()
	resp = make([]queueResponse, 0, Limit)
	b.Call("queue", []int{LastId, Limit}, &resp)
	return
}
