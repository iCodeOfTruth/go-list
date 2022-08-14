package golist

import (
	"math/rand"
	"sync"
)

type ListManager struct {
	List []interface{}
	lock *sync.Mutex

	currentIndex int
}

func NewListManager(list []interface{}) *ListManager {
	return &ListManager{List: list, lock: &sync.Mutex{}}
}

func (listM *ListManager) Random() (element interface{}) {
	listM.lock.Lock()
	defer listM.lock.Unlock()
	return listM.List[rand.Intn(len(listM.List))]
}

func (listM *ListManager) Next() (element interface{}) {
	listM.lock.Lock()
	defer listM.lock.Unlock()
	if listM.currentIndex >= len(listM.List) {
		listM.currentIndex = 0
	}

	listM.currentIndex++
	return listM.List[listM.currentIndex-1]
}

func (listM *ListManager) RemoveIndex(index int) {
	listM.lock.Lock()
	defer listM.lock.Unlock()
	listM.List = append(listM.List[:index], listM.List[index+1:]...)
}

func (listM *ListManager) Remove(item interface{}) {
	for i, v := range listM.List {
		if v == item {
			listM.RemoveIndex(i)
			return
		}
	}
}

func (listM *ListManager) Append(item interface{}) {
	listM.lock.Lock()
	defer listM.lock.Unlock()
	listM.List = append(listM.List, item)
}

func (listM *ListManager) Extend(items []interface{}) {
	listM.lock.Lock()
	defer listM.lock.Unlock()
	listM.List = append(listM.List, items...)
}
