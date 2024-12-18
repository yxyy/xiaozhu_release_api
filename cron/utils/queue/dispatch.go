package queue

import (
	"sync"
)

var list []*Queue
var mu sync.Mutex

func AddQueue(q ...*Queue) {
	mu.Lock()
	defer mu.Unlock()
	list = append(list, q...)
}

func Run() {
	for _, v := range list {
		if v != nil {
			go v.Run()
		}
	}
}
