package concurrency

import "sync"

type Collection struct {
	Mutex sync.Mutex
	Data  map[string]string
}

func NewCollection() Collection {
	return Collection{
		Mutex: sync.Mutex{},
		Data:  make(map[string]string),
	}
}

func (c *Collection) Add(key, value string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if c.Has(key) {
		return
	}
	c.Data[key] = value
}

func (c *Collection) Has(key string) bool {
	_, ok := c.Data[key]
	return ok
}
