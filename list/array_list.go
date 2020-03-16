package list

import (
	"fmt"

	"github.com/SsrCoder/go-data-structure/errors"
	"github.com/SsrCoder/go-data-structure/utils"
)

// arrayList 动态数组
type arrayList struct {
	arr []interface{}
}

// NewArrayList 创建ArrayList对象
func NewArrayList() *arrayList {
	return &arrayList{}
}

// Len 获取列表尺寸
func (l *arrayList) Len() int {
	return len(l.arr)
}

// Empty 判断列表是否为空
func (l *arrayList) Empty() bool {
	return l.Len() == 0
}

// Contains 判断列表中是否包含指定元素
func (l *arrayList) Contains(val interface{}) bool {
	return l.IndexOf(val) != -1
}

// Add 在列表的尾部添加元素
func (l *arrayList) Add(val interface{}) {
	l.arr = append(l.arr, val)
}

// Get 获取指定索引位置的元素
func (l *arrayList) Get(idx int, ret interface{}) error {
	if idx < 0 || idx >= l.Len() {
		return fmt.Errorf(`%w: {"idx": %d, "length": %d}`, errors.ErrIndexOutOfBounds, idx, l.Len())
	}

	return utils.Set(l.arr[idx], ret)
}

// Set 修改置顶索引的元素
func (l *arrayList) Set(idx int, val interface{}) error {
	if idx < 0 || idx > l.Len() {
		return fmt.Errorf(`%w: {"idx": %d, "length": %d}`, errors.ErrIndexOutOfBounds, idx, l.Len())
	} else if idx == l.Len() {
		l.Add(val)
	} else {
		l.arr[idx] = val
	}
	return nil
}

// Insert 在指定索引位置插入元素
func (l *arrayList) Insert(idx int, val interface{}) error {
	if idx < 0 || idx > l.Len() {
		return fmt.Errorf(`%w: {"idx": %d, "length": %d}`, errors.ErrIndexOutOfBounds, idx, l.Len())
	} else if idx == l.Len() {
		l.Add(val)
	} else {
		l.arr = append(l.arr[:idx], append([]interface{}{val}, l.arr[idx:]...))
	}
	return nil
}

// Remove 移除置顶索引位置的元素
func (l *arrayList) Remove(idx int) error {
	if idx < 0 || idx >= l.Len() {
		return fmt.Errorf(`%w: {"idx": %d, "length": %d}`, errors.ErrIndexOutOfBounds, idx, l.Len())
	} else if idx == l.Len()-1 {
		l.arr = l.arr[:idx]
	} else {
		l.arr = append(l.arr[:idx], l.arr[idx+1:])
	}
	return nil
}

// IndexOf 获取指定元素在列表中的索引，如果不存在返回 -1
func (l *arrayList) IndexOf(val interface{}) int {
	for i, v := range l.arr {
		if v == val {
			return i
		}
	}
	return -1
}

// Clear 清空列表
func (l *arrayList) Clear() {
	l.arr = make([]interface{}, 0)
}

// ForEach 遍历数组的每个元素
func (l *arrayList) ForEach(walk func(val interface{})) {
	for _, v := range l.arr {
		walk(v)
	}
}
