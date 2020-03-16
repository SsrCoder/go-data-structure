package utils

import (
	"reflect"

	"github.com/SsrCoder/go-data-structure/errors"
)

// Set 将 val 的值传递到ret中去，其中ret必须是指针类型
func Set(val interface{}, ret interface{}) error {
	if reflect.TypeOf(ret).Kind() != reflect.Ptr {
		return errors.ErrTypeNotPointer
	}
	reflect.ValueOf(ret).Elem().Set(reflect.ValueOf(val))
	return nil
}
