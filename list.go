package ulist

import "sync"

type Item interface {
	Equal(Item) bool
}

type List struct {
	Entities []Item
	mu       sync.RWMutex
}

func New(list []Item) *List {
	return &List{
		Entities: list,
	}
}

func (l *List) Add(v Item) {
	if l.Has(v) {
		return
	}

	l.mu.Lock()
	l.Entities = append(l.Entities, v)
	l.mu.Unlock()
}

func (l *List) Has(v Item) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for _, e := range l.Entities {
		if e.Equal(v) {
			return true
		}
	}

	return false
}

func (l *List) Index(v Item) int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for i, e := range l.Entities {
		if e.Equal(v) {
			return i
		}
	}

	return -1
}

func (l *List) Get(v Item) Item {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for _, e := range l.Entities {
		if e.Equal(v) {
			return e
		}
	}

	return nil
}

func (l *List) All() []Item {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.Entities
}

func (l *List) Del(v Item) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for i, e := range l.Entities {
		if e.Equal(v) {
			l.mu.Lock()
			l.Entities = append(l.Entities[:i], l.Entities[i+1:]...)
			l.mu.Unlock()
			break
		}
	}
}
