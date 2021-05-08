package geek_demo

import (
	"context"
	"sync/atomic"
	"time"
)

type TaskHandler func() error

type WPool struct {
	errChan chan error
	task    chan TaskHandler
	temTask chan TaskHandler
	done    chan bool
	last    int64
	pool    int64
	timeout time.Duration
}

func (p WPool) start() {
	for fn := range p.task {
		ctx, cancel := context.WithTimeout(context.TODO(), p.timeout)
		closed := make(chan struct{}, 1)
		go func(fn TaskHandler) {
			err := fn()
			if err != nil {
				p.errChan <- err
			}
			close(closed)
		}(fn)

		go func() {
			select {
			case <-ctx.Done():
				p.done <- true
			case <-closed:
				p.done <- true
				cancel()
			}
		}()
	}
}

func NewPool(max, count int) *WPool {
	if max < 1 {
		max = 1
	}

	p := &WPool{
		errChan: make(chan error, max),
		task:    make(chan TaskHandler, max),
		temTask: make(chan TaskHandler, max),
		done:    make(chan bool, max),
		last:    int64(max),
		pool:    int64(count),
		timeout: 5 * time.Second,
	}

	go func() {
		for {
			<-p.done
			lastPool := atomic.AddInt64(&p.pool, 1)
			atomic.AddInt64(&p.last, -1)
			if lastPool > 0 && len(p.temTask) > 0 {
				fn := <-p.temTask
				p.task <- fn
				atomic.AddInt64(&p.pool, -1)
			}
		}
	}()

	go p.start()

	return p
}

func (p WPool) Do(fn TaskHandler) {
	lastPool := atomic.LoadInt64(&p.pool)
	if lastPool > 0 {
		atomic.AddInt64(&p.pool, -1)
		p.task <- fn
	} else {
		p.temTask <- fn
	}
}

func (p WPool) SetTimeout(timeout time.Duration) {
	p.timeout = timeout
}
