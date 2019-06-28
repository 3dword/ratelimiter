package bucket

import (
	"sync"
)

type Bucket struct {
	allowance uint64
	capacity uint64
	lastDuration int64
	mutex sync.Mutex
}

func (b *Bucket) SetLastDuration(lastDuration int64) {
	b.mutex.Lock()
	defer func() {
		b.mutex.Unlock()
	}()

	b.lastDuration = lastDuration
}

func (b *Bucket) Tick() {
	b.mutex.Lock()
	defer func() {
		b.mutex.Unlock()
	}()
	b.capacity--
}

func (b *Bucket) Reset() {
	b.mutex.Lock()
	defer func() {
		b.mutex.Unlock()
	}()
	b.capacity = b.allowance
}

func (b *Bucket) SetCapacity(capacity uint64) {
	b.mutex.Lock()
	defer func() {
		b.mutex.Unlock()
	}()
	b.capacity = capacity
}

func (b *Bucket) GetCapacity() uint64{
	return b.capacity
}

var b *Bucket

func New(allowance uint64, lastDuration int64) *Bucket{
	if b != nil {
		return b
	}
	b = &Bucket{
		allowance : allowance,
		lastDuration : lastDuration,
		capacity: allowance,
	}

	return b
}