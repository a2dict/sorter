package sorter

import (
	"fmt"
	"reflect"
	"sort"
)

// Comparator ...
type Comparator func(a, b interface{}) int

// Predicate ...
type Predicate func(a interface{}) bool

// Extractor ...
type Extractor func(a interface{}) interface{}

type sorter struct {
	data interface{}
	cmpr Comparator
}

// Len ...
func (s sorter) Len() int {
	return reflect.ValueOf(s.data).Elem().Len()
}

// Less ...
func (s sorter) Less(i, j int) bool {
	arr := reflect.ValueOf(s.data).Elem()
	a := arr.Index(i).Interface()
	b := arr.Index(j).Interface()
	res := s.cmpr(a, b)
	if res < 0 {
		return true
	}
	return false
}

// Swap ...
func (s sorter) Swap(i, j int) {
	if i > j {
		i, j = j, i
	}
	arr := reflect.ValueOf(s.data).Elem()

	tmp := arr.Index(i).Interface()
	arr.Index(i).Set(arr.Index(j))
	arr.Index(j).Set(reflect.ValueOf(tmp))
}

// NewSorter ...
func NewSorter() *sorter {
	return &sorter{}
}

// Comparing ...
func (s *sorter) Comparing(comparator Comparator) *sorter {
	if s.cmpr == nil {
		return &sorter{cmpr: comparator}
	}
	return &sorter{cmpr: func(a, b interface{}) int {
		res := s.cmpr(a, b)
		if res != 0 {
			return res
		} else {
			return comparator(a, b)
		}
	}}
}

// ComparingBy ...
func (s *sorter) ComparingBy(extractor Extractor) *sorter {
	return s.Comparing(extractor.toComparator())
}

// ReversedComparing ...
func (s *sorter) ReversedComparing(comparator Comparator) *sorter {
	return s.Comparing(comparator.flip())
}

// ReversedComparingBy ...
func (s *sorter) ReversedComparingBy(extractor Extractor) *sorter {
	return s.Comparing(extractor.toComparator().flip())
}

// MoveForward ...
func (s *sorter) MoveForward(predicate Predicate) *sorter {
	return s.ComparingBy(func(a interface{}) interface{} {
		return !predicate(a)
	})
}

// MoveBackward ...
func (s *sorter) MoveBackward(predicate Predicate) *sorter {
	return s.ComparingBy(func(a interface{}) interface{} {
		return predicate(a)
	})
}

// Sort ...
func (s *sorter) Sort(data interface{}) {

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Ptr ||
		v.IsNil() ||
		v.Elem().Kind() != reflect.Slice {
		panic("data should be array pointer(*[])")
	}

	st := sorter{data: data, cmpr: s.cmpr}
	sort.Sort(st)
}

func (c Comparator) flip() Comparator {
	return func(a, b interface{}) int {
		return c(b, a)
	}
}

func (e Extractor) toComparator() Comparator {
	return func(a, b interface{}) int {
		ea := e(a)
		eb := e(b)
		return ordering(ea, eb)
	}
}

// prime data order
func ordering(a, b interface{}) int {
	if a == b {
		return 0
	}
	switch a.(type) {
	case string:
		return lg(a.(string) > b.(string))
	case int:
		return lg(a.(int) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int))
	case int8:
		return lg(a.(int8) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int8))
	case int16:
		return lg(a.(int16) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int16))
	case int32:
		return lg(a.(int32) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int32))
	case int64:
		return lg(a.(int64) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int64))
	case uint:
		return lg(a.(uint) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint))
	case uint8:
		return lg(a.(uint8) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint8))
	case uint16:
		return lg(a.(uint16) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint16))
	case uint32:
		return lg(a.(uint32) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint32))
	case uint64:
		return lg(a.(uint64) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint64))
	case float32:
		return lg(a.(float32) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(float32))
	case float64:
		return lg(a.(float64) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(float64))
	case bool:
		return lg(a.(bool) && !b.(bool))
	default:
		panic(fmt.Sprintf("dont know how to compare: %T", a))
	}
}

func lg(b bool) int {
	if b {
		return 1
	}
	return -1
}
