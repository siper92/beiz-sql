package beiz_sql

import (
	core_utils "github.com/siper92/core-utils"
	"reflect"
)

type MVal[T any] struct {
	val T
}

func (m MVal[T]) IsNil() bool {
	return m.val == nil
}

//func (m MaybeVal) Get() interface{} {
//	return m.val
//}

func GetStructFieldOrNil(e any, s string) *reflect.StructField {
	t := reflect.TypeOf(e)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil
	}

	f, ok := t.FieldByName(s)
	if !ok {
		core_utils.Warning("Field %s not found in struct %s", s, t.Name())
		return nil
	}

	return &f
}
