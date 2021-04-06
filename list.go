package ulist

import "sync"

type Item interface {
	Equal(Item) bool
	Merge(Item)
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
	x := l.Get(v)

	if x != nil {
		x.Merge(v)
	} else {
		l.mu.Lock()
		l.Entities = append(l.Entities, v)
		l.mu.Unlock()
	}
}

func (l *List) Has(v Item) bool {
	return l.Index(v) >= 0
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

	for i := 0; i < len(l.Entities); i++ {
		if l.Entities[i].Equal(v) {
			l.mu.Lock()
			l.Entities[i] = l.Entities[0]
			l.Entities = l.Entities[1:]
			l.mu.Unlock()
			i--
		}
	}
}

func (l *List) Merge() {
	l.mu.Lock()
	defer l.mu.Unlock()

	for ai := 0; ai < len(l.Entities); ai++ {
		a := l.Entities[ai]
		for bi := 0; bi < len(l.Entities); bi++ {
			b := l.Entities[bi]
			if a.Equal(b) {
				a.Merge(b)

				l.Entities[bi] = l.Entities[0]
				l.Entities = l.Entities[1:]
				i--
			}
		}
	}
}
