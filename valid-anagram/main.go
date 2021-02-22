package main

func isAnagram(s string, t string) bool {
	counter := map[string]int{}

	if len(s) != len(t) {
		return false
	}

	for _, val := range s {
		character := string(val)
		if _, ok := counter[character]; !ok {
			counter[character] = 1
		} else {
			counter[character] += 1
		}
	}

	for _, val := range t {
		character := string(val)
		if v, ok := counter[character]; ok {
			if v-1 == 0 {
				delete(counter, character)
			} else {
				counter[character] = v - 1
			}
		}
	}
	return len(counter) == 0
}
