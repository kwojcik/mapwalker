package nogrowmap

import "math/rand"

// NoGrowMap simulates a map that can be iterated over and inserted into, but never has to grow
type NoGrowMap struct {
	Iterator chan struct{} // Iterator represents all the elements left to go until the end of the map
	Size     int           // Size represents the total size of the map
}

// NewNoGrowMap allocates a NoGrowMap with some initialSize
func NewNoGrowMap(initialSize int) *NoGrowMap {
	// make an iterator that has initialSize items left to go
	iterator := make(chan struct{}, initialSize)
	for i := 0; i < initialSize; i++ {
		iterator <- struct{}{}
	}
	return &NoGrowMap{
		Size:     initialSize,
		Iterator: iterator,
	}
}

// Insert simulates inserting into a NoGrowMap
func (m *NoGrowMap) Insert() {
	if rand.Intn(m.Size+1) > m.iteratorPos() {
		// Something got inserted infront of the iterator's position in the map
		m.Iterator <- struct{}{}
	}
	m.Size++

	if len(m.Iterator) == 0 {
		// We have reached the end!
		close(m.Iterator)
	}
}

// iteratorPos gives the iterators current position in the map's array layout
func (m *NoGrowMap) iteratorPos() int {
	return m.Size - len(m.Iterator)
}
