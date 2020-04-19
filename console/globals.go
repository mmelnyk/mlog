package console

import (
	"sync"
)

var (
	eventPool *sync.Pool
)

func init() {
	eventPool = &sync.Pool{
		New: func() interface{} {
			evt := &event{}
			return evt
		},
	}
}

func getEvent() *event {
	return eventPool.Get().(*event)
}

func putEvent(evt *event) {
	if evt != nil && evt.buffer.Len() < maxAllowedEventBuffer {
		evt.buffer.Reset()
		eventPool.Put(evt)
	}
}
