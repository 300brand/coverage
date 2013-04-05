package bridge

func Queue() (out interface{}, err error) {
	/*
		c, err := Client()
		if err != nil {
			return
		}
	*/
	c := Client()
	defer c.Close()
	c.Call("queue", 0, &out)
	return
}
