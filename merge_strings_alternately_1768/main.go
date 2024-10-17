package main

func mergeAlternately(word1 string, word2 string) string {
	acc := ""
	for idx := range max(len(word1), len(word2)) {
		if idx < len(word1) {
			acc += string(word1[idx])
		}
		if idx < len(word2) {
			acc += string(word2[idx])
		}
	}
	return acc
}

func maxNumber(first int, second int) int {
	if first > second {
		return first
	}
	return second
}

func main() {
}
