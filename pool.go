package pool

import (
	"sync"
)

// A Pool is a generic wrapper around a sync.Pool.
type Pool[T any] struct {
	pool sync.Pool
}

// NewPool creates a new Pool with the provided new function.
//
// The equivalent sync.Pool construct is "sync.Pool{New: fn}"
func NewPool[T any](fn func() T) Pool[T] {
	return Pool[T]{
		pool: sync.Pool{New: func() any { return fn() }},
	}
}

// Get is a generic wrapper around sync.Pool's Get method.
func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}

// Put is a generic wrapper around sync.Pool's Put method.
func (p *Pool[T]) Put(x T) {
	p.pool.Put(x)
}

type SlicePool[T any] struct {
	syncSlicePool sync.Pool
}

func NewSlicePool[T any](size int) *SlicePool[T] {
	return &SlicePool[T]{
		syncSlicePool: sync.Pool{New: func() any {
			return make([]T, 0, size)
		}},
	}
}

func (p *SlicePool[T]) Get() []T {
	return p.syncSlicePool.Get().([]T)
}

func (p *SlicePool[T]) Put(s []T) {
	p.syncSlicePool.Put(s[:0])
}
