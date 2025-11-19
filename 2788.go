package main

func splitWordsBySeparator(words []string, separator byte) []string {
	var ret []string
	add := func(b []byte) {
		if len(b) == 0 {
			return
		}
		ret = append(ret, string(b))
	}
	for _, w := range words {
		var b []byte
		for _, c := range w {
			if byte(c) != separator {
				b = append(b, byte(c))
				continue
			}
			add(b)
			b = b[:0]
		}
		add(b)
	}
	return ret
}
