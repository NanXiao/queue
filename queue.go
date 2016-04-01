package queue

type Queue struct {
	queue []interface{}
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Len() int {
	return len(q.queue)
}

func (q *Queue) Enqueue(v interface{})  {
	q.queue = append(q.queue, v)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.queue) == 0 {
		return nil
	} else {
		v := q.queue[0]
		q.queue = q.queue[1:]
		return v
	}
}

func (q *Queue) Head() interface{} {
	if len(q.queue) == 0 {
		return nil
	} else {
		return q.queue[0]
	}
}