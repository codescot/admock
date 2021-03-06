package cache

import (
	"sort"
	"sync"

	"github.com/codescot/go-common/array"
	"github.com/codescot/go-common/math"
)

// StringCache simple sorted key cache
type StringCache struct {
	mux  sync.Mutex
	data []string
	Size int
}

// Strings create new string cache
func Strings() *StringCache {
	return &StringCache{
		mux:  sync.Mutex{},
		data: []string{},
	}
}

// Get get a value
func (sc *StringCache) Get(i int) string {
	return sc.data[i]
}

// Add add a value
func (sc *StringCache) Add(value string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()

	sc.data = append(sc.data, value)
	sc.Size = len(sc.data)
}

// Append append a list of values
func (sc *StringCache) Append(values []string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()

	sc.data = append(sc.data, values...)
	sc.Size = len(sc.data)
}

// Remove remove a value
func (sc *StringCache) Remove(value string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()

	i := sort.SearchStrings(sc.data, value)
	sc.data = append(sc.data[:i], sc.data[i+1:]...)
	sc.Size = len(sc.data)
}

// Sort sort the cache
func (sc *StringCache) Sort() {
	sc.mux.Lock()
	defer sc.mux.Unlock()

	if !sort.StringsAreSorted(sc.data) {
		sort.Strings(sc.data)
	}
}

// Contains check if cache contains a value
func (sc *StringCache) Contains(value string) bool {
	sc.mux.Lock()
	defer sc.mux.Unlock()

	i := sort.SearchStrings(sc.data, value)

	start := math.Max(i-1, 0)
	end := math.Min(i+1, sc.Size)

	s := sc.data[start:end]

	return array.Contains(s, value)
}
