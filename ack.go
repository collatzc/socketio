package socketio

import (
	"errors"
	"sync"
)

var (
	ErrorWaiterNotFound = errors.New("Waiter not found")
)

type ackProcessor struct {
	counter          int
	counterLock      sync.Mutex
	resultWaitersMap sync.Map
}

func (a *ackProcessor) getNextId() int {
	a.counterLock.Lock()
	defer a.counterLock.Unlock()

	a.counter++
	return a.counter
}

func (a *ackProcessor) addWaiter(id int, w chan interface{}) {
	a.resultWaitersMap.Store(id, w)
}

func (a *ackProcessor) removeWaiter(id int) {
	a.resultWaitersMap.Delete(id)
}

func (a *ackProcessor) getWaiter(id int) (chan interface{}, error) {
	if waiter, ok := a.resultWaitersMap.Load(id); ok {
		return waiter.(chan interface{}), nil
	}
	return nil, ErrorWaiterNotFound
}
