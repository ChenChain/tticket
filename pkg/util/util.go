package util

import (
	"math/rand"
	"time"
)

type ValueType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
}

// ArrayDifferentSet 数组求差集
func ArrayDifferentSet[T ValueType](arr1, arr2 []T) []T {
	res := make([]T, 0)
	arrMap := make(map[T]bool)
	for _, v := range arr2 {
		arrMap[v] = true
	}
	for _, v := range arr1 {
		if !arrMap[v] {
			res = append(res, v)
		}
	}
	return res
}

func ArrayIntersect[T ValueType](arr1, arr2 []T) []T {
	res := make([]T, 0)
	map1 := make(map[T]bool)
	for _, v1 := range arr1 {
		map1[v1] = true
	}
	for _, v2 := range arr2 {
		if map1[v2] {
			res = append(res, v2)
		}
	}
	return res
}

// ArrayAdd 数组相加， 不去重
func ArrayAdd[T ValueType](arr1, arr2 []T) []T {
	for _, v := range arr2 {
		arr1 = append(arr1, v)
	}
	return arr1
}

func RandomChooseOne[T ValueType](arr []T) T {
	length := len(arr)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := r.Intn(length)
	return arr[res]
}
