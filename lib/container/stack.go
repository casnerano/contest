package container

type Stack[T any] struct {
    items []*T
}

func (q *Stack[T]) Push(item *T) {
    q.items = append(q.items, item)
}

func (q *Stack[T]) Pop() (*T, error) {
    if q.Empty() {
        return nil, ErrEmptyContainer
    }

    lastIndex := q.Size() - 1
    item := q.items[lastIndex]
    q.items = q.items[:lastIndex]

    return item, nil
}

func (q *Stack[T]) Peek() (*T, error) {
    if q.Empty() {
        return nil, ErrEmptyContainer
    }

    lastIndex := q.Size() - 1

    return q.items[lastIndex], nil
}

func (q *Stack[T]) Size() int {
    return len(q.items)
}

func (q *Stack[T]) Empty() bool {
    return q.Size() == 0
}

func (q *Stack[T]) Clear() {
    q.items = nil
}
