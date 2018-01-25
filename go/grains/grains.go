package grains

import "errors"

var valueMap map[int]uint64 = make(map[int]uint64)

func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("n is NOT in range")
	}
	if valueMap[n] > 0 {
		return valueMap[n], nil
	}
	var r uint64 = 0
	if n == 1 {
		r = 1
	} else {
		prev, _ := Square(n - 1)
		r = prev * 2
	}
	valueMap[n] = r
	return r, nil
}

func Total() uint64 {
	var sum uint64 = 0
	for i := 1; i <= 64; i++ {
		if !(valueMap[i] >= 0) {
			Square(i)
		}
		sum += valueMap[i]
	}
	return sum
}
