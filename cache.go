package expirable

import (
	"sync"
	"time"
)

type Operation int

const (
	Noop Operation = 1 << iota
	Create
	Read
	Update
	Delete
)

type Keyable interface{}

type Expirable interface {
	Expired(baseTime time.Time) bool
	Value() interface{}
}

type OperationCache struct {
	mu   *sync.Mutex
	body map[hashItem]hashItem
}

func (oc *OperationCache) Put(k Keyable, v Expirable) bool {
	oc.mu.Lock()
	defer oc.mu.Unlock()

	oc.body[k] = v
}

func (oc *OperationCache) Get(k Keyable) (Expirable, bool) {
	oc.mu.Lock()
	defer oc.mu.Unlock()

	v, ok := oc.body[k]
	return v, ok
}
