package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(strs []string) FreqMap {
	m := FreqMap{}
	for _, s := range strs {
		single_fm := Frequency(s)
		for l := range single_fm {
			m[l] += single_fm[l]
		}
	}
	return m
}
