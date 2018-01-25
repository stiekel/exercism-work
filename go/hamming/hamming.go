package hamming

import "errors"

func Distance(a, b string) (int, error) {
	d := 0
	if len(a) != len(b) {
		return d, errors.New("not equal")
	}

	for i := range a {
		if a[i] != b[i] {
			d += 1
		}
	}
	return d, nil
}
