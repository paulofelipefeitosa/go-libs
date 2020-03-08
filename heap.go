package heap

type Heap struct {
	values     []int
	last       int
	comparator func(int, int) bool
}

func New(size int, comp func(int, int) bool) (*Heap, error) {
	if size < 0 {
		return nil, fmt.Errorf("Heap size cannot be negative, given size value (%d)", size)
	}
	if comp == nil {
		return nil, fmt.Errorf("Heap comparator function cannot be nil")
	}
	return &Heap{values: make([]int, size, size), last: 0, comparator: comp}, nil
}

func (heap *Heap) Insert(value int) {
	if heap == nil {
		panic("Heap is nil")
	}
	if len(heap.values) == heap.last {
		newSize := maxInt(2, len(heap.values)<<1)
		newValues := make([]int, newSize, newSize)
		copy(newValues, heap.values)
		heap.values = newValues
	}
	heap.values[heap.last] = value
	heap.upHeapify(heap.last)
	heap.last++
}

func (heap *Heap) upHeapify(startNode int) {
	for node := startNode; node != 0; node = parent(node) {
		parent := parent(node)
		if heap.compare(node, parent) {
			heap.swap(node, parent)
		}
	}
}

func (heap *Heap) ExtractTop() (int, bool) {
	if heap == nil {
		panic("Heap is nil")
	}
	if heap.Empty() {
		return 0, false
	}
	top := heap.values[0]
	heap.swap(0, heap.last-1)
	heap.last--
	heap.heapify(0)
	return top, true
}

func (heap *Heap) heapify(node int) {
	if node >= heap.last {
		return
	}
	left := left(node)
	right := right(node)
	if left >= heap.last && right >= heap.last {
		return
	}
	
	better := left
	if left < heap.last && right < heap.last {
		if heap.compare(right, left) {
			better = right
		}
	} else if right < heap.last {
		better = right
	}
	
	if heap.compare(better, node) {	
		heap.swap(node, better)
	}
	heap.heapify(better)
}

func (heap *Heap) Empty() bool {
	if heap == nil {
		panic("Heap is nil")
	}
	if heap.last == 0 {
		return true
	}
	return false
}

func (heap *Heap) compare(node, otherNode int) bool {
	return heap.comparator(heap.values[node], heap.values[otherNode])
}

func (heap *Heap) swap(node, otherNode int) {
	heap.values[node], heap.values[otherNode] = heap.values[otherNode], heap.values[node]
}

// helper functions below

func parent(node int) int {
	return (node - 1) >> 1
}

func left(node int) int {
	return (node << 1) + 1
}

func right(node int) int {
	return (node << 1) + 2
}

func maxInt(x, y int) int {
	if y > x {
		return y
	}
	return x
}
