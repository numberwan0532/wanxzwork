package main

// 最长前缀
func getLongestPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	byteStrs := []byte{}
	for i := 0; i < len(strs[0]); i++ {
		falg := false
		marr := make(map[string]int)
		for _, v := range strs {
			if i >= len(v) {
				falg = true
				break
			}
			marr[string(v[i])]++
		}
		if falg {
			break
		}
		if len(marr) == 1 {
			for k := range marr {
				byteStrs = append(byteStrs, k[0])
			}
		} else {
			falg = true
			break
		}
		if falg {
			break
		}

	}
	return string(byteStrs)
}
