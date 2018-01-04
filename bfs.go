package algorithm

type BFS struct {
	queue        []interface{}
	head, tail   int
	depth        int
	currentLevel int
	nextLevel    int
}

func NewBFS() *BFS {
	b := new(BFS)
	b.queue = nil
	b.head = 0
	b.tail = 0
	b.depth = 0
	b.currentLevel = 0
	b.nextLevel = 0
	return b
}

func (b *BFS) Push(v interface{}) {
	b.queue = append(b.queue, v)
	b.head++
	b.nextLevel++
}

func (b *BFS) Pop() interface{} {
	if b.tail >= b.head {
		return nil
	}
	v := b.queue[b.tail]
	b.tail++
	if b.currentLevel <= 0 {
		b.depth++
		b.currentLevel = b.nextLevel
		b.nextLevel = 0
	}
	b.currentLevel -= 1
	return v
}
