func mergeAlternately(word1 string, word2 string) string {
	word1Runes := []rune(word1)
	word2Runes := []rune(word2)
	var result []string

	for i := 0; i < max(len(word2Runes), len(word1Runes)); i++ {
		w2 := stringAt(word2, i)
        w1 := stringAt(word1, i)

		result = append(result, w1, w2)
	}

	answer := strings.Join(result, "")
    return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stringAt(s string, index int) string {
	if index < len(s) {
		return string(s[index])
	}
	return ""
}