package utils

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	Capacity int
	Cache    map[int]*list.Element
	List     *list.List
	Mu       sync.Mutex
}

type Pair struct {
	Key   int
	Value int
}

func CacheConstructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*list.Element),
		List:     list.New(),
		Mu:       sync.Mutex{},
	}
}

func (lru *LRUCache) Get(key int) int {
	lru.Mu.Lock()
	defer lru.Mu.Unlock()

	if element, exists := lru.Cache[key]; exists {
		lru.List.MoveToFront(element)
		return element.Value.(Pair).Value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	lru.Mu.Lock()
	defer lru.Mu.Unlock()

	if element, exists := lru.Cache[key]; exists {
		lru.List.MoveToFront(element)
		element.Value = Pair{key, value}
	}

	if lru.List.Len() == lru.Capacity {
		delete(lru.Cache, lru.List.Back().Value.(Pair).Key)
		lru.List.Remove(lru.List.Back())
	}

	element := lru.List.PushFront(Pair{key, value})
	lru.Cache[key] = element
}
