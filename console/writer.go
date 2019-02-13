package console

import (
	"io"
	"sync"
)

// syncwriter is wrapper around Writer interface with mutex lock
type syncwriter struct {
	w  io.Writer
	mu sync.Mutex
}

func (w *syncwriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	n, err = w.w.Write(p)
	w.mu.Unlock()
	return
}
