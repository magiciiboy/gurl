package common

import (
	"fmt"
	"sort"
)

// MinInt64NonZero returns min value of 2 ints
func MinInt64NonZero(x, y int64) int64 {
	if x == 0 || y > x {
		return y
	}
	return x
}

// MaxInt64 returns max value of 2 ints
func MaxInt64(x, y int64) int64 {
	if y > x {
		return y
	}
	return x
}

// GetMin returns min of an array
func GetMin(arr []int) (min int, err error) {
	if arr == nil || len(arr) == 0 {
		err = fmt.Errorf("Empty array")
		return
	}
	for i, n := range arr {
		if n < min || i == 0 {
			min = n
		}
	}
	return
}

// GetMax returns max of an array
func GetMax(arr []int) (max int, err error) {
	if arr == nil || len(arr) == 0 {
		err = fmt.Errorf("Empty array")
		return
	}
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	return
}

// GetMedian returns median of an array
func GetMedian(arr []int) (median float32, err error) {
	sort.Ints(arr)
	l := len(arr)
	if arr == nil || l == 0 {
		err = fmt.Errorf("Empty array")
		return
	}
	if l%2 == 0 {
		return (float32(arr[l/2]) + float32(arr[(l-1)/2])) / 2, nil
	}
	return float32(arr[l/2]), nil
}
