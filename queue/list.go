package queue

//first in first out

const capacity = 100

type Queue struct {
	size  int
	data  [capacity]int
	front int
	back  int
}

func (q *Queue) Add(value int) bool {
	if q.size >= capacity {
		return false
	}
	q.size++
	q.data[q.back] = value
	q.back = (q.back + 1) % (capacity - 1)
	return true
}

func (q *Queue) Remove() (int, bool) {
	var value int
	if q.size <= 0 {
		return 0, false
	}
	q.size--
	value = q.data[q.front]
	q.front = (q.front + 1) % (capacity - 1)
	return value, true
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}
