package testfixture

import (
	"fmt"
	"sync"
)

var mappings sync.Map

func store(name string, prefix string, key any, val any) {
	m, _ := mappings.LoadOrStore(name, &sync.Map{})
	m.(*sync.Map).Store(fmt.Sprintf("%s:%v", prefix, key), val)
}

func Get[T any](name string, prefix string, key any) T {
	var t T
	m, ok := mappings.Load(name)
	if !ok {
		return t
	}
	v, ok := m.(*sync.Map).Load(fmt.Sprintf("%s:%v", prefix, key))
	if !ok {
		return t
	}
	return v.(T)
}
