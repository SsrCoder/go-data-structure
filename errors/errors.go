package errors

import (
	"errors"
)

var (
	// ErrIndexOutOfBounds 数组下标越界
	ErrIndexOutOfBounds = errors.New("index out of bounds")
	// ErrTypeNotPointer 对象非指针类型
	ErrTypeNotPointer = errors.New("not pointer")
)
