package container

import "errors"

var ErrEmptyContainer = errors.New("empty container")

type Container interface {
    Size() int
    Empty() bool
    Clear()
}
