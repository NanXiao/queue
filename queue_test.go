package queue

import "testing"

func TestNewQueue(t *testing.T)  {
	q := New()
	if q == nil {
		t.Error("Queue must not be nil")
	}
}

func TestEmptyQueue(t *testing.T)  {
	var q Queue
	if q.Len() != 0 || q.Dequeue() != nil || q.Head() != nil{
		t.Error("Queue must have no element")
	}
}

func TestEnqueue(t *testing.T)  {
	q := New()

	q.Enqueue(1)
	if q.Len() != 1 {
		t.Error("Queue must have one element")
	}
	if q.Head().(int) != 1 {
		t.Error("The first element must be 1")
	}
}

func TestDequeue(t *testing.T)  {
	q := New()

	q.Enqueue(1)
	q.Enqueue(2)
	if q.Len() != 2 {
		t.Error("Queue must have two elements")
	}

	v := q.Dequeue()
	if v.(int) != 1 {
		t.Error("The v must be 1")
	}
	if q.Len() != 1 {
		t.Error("Queue must have one element")
	}
	if q.Head().(int) != 2 {
		t.Error("The head be 2")
	}
}