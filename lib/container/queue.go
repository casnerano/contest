package container

type Queue[T any] struct {
    items []*T
}

func (q *Queue[T]) Push(item *T) {
    q.items = append(q.items, item)
}

func (q *Queue[T]) Pop() (*T, error) {
    if q.Empty() {
        return nil, ErrEmptyContainer
    }

    item := q.items[0]
    q.items = q.items[1:]

    return item, nil
}

func (q *Queue[T]) Peek() (*T, error) {
    if q.Empty() {
        return nil, ErrEmptyContainer
    }

    return q.items[0], nil
}

func (q *Queue[T]) Size() int {
    return len(q.items)
}

func (q *Queue[T]) Empty() bool {
    return q.Size() == 0
}

func (q *Queue[T]) Clear() {
    q.items = nil
}
