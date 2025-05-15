package patterns

// ensures a class has only one instance

import (
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
	callCounter int
}

func (s *single) getCallCounter() int {
	return s.callCounter
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{
				callCounter: 1,
			}
		}
	} else {
		singleInstance.callCounter++
	}

	return singleInstance
}
