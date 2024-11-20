package utils

import (
	"xiaozhu/backend/internal/model/common"
)

type Number interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64 | uint | ~int
}

type Slice interface {
	Number | string | byte | rune
}

func ArrayUnique[T Slice | int32](nums []T) []T {
	if len(nums) == 0 {
		return nums
	}

	var mp = make(map[T]int)
	for _, v := range nums {
		mp[v] = 1
	}

	return MapKeys(mp)
}

func MapKeys[T Slice, V comparable](data map[T]V) []T {
	if data == nil {
		return nil
	}
	var list []T
	for k := range data {
		list = append(list, k)
	}

	return list
}

func MapValue[T Slice](data map[T]T) []T {
	if data == nil {
		return nil
	}
	var list []T
	for _, v := range data {
		list = append(list, v)
	}

	return list
}

func ConvertIdNameMapById(list []*common.IdName) map[int]*common.IdName {
	result := make(map[int]*common.IdName)
	for _, v := range list {
		result[v.Id] = v
	}

	return result
}
