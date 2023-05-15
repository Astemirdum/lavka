package pool

import "sync"

type TimePool struct {
	pool *sync.Pool
}

const (
	dayInMinute = 24*60 + 1
)

type Buf [dayInMinute]int

func NewTimePool() *TimePool {
	return &TimePool{pool: &sync.Pool{New: func() any {
		return new(Buf)
	}}}
}

func (p *TimePool) Get() *Buf {
	return p.pool.Get().(*Buf)
}

func (p *TimePool) Put(buf *Buf) {
	for i := range buf {
		buf[i] = 0
	}
	p.pool.Put(buf)
}
