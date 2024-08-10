package utils

func toNum(c byte) int {
	nums := "0123456789"
	for i := range nums {
		if c == nums[i] {
			return i
		}
	}
	return -1
}

func ExtractNums(s string) []int {
	result := []int{}
	i := 0

	for i < len(s) {
		for toNum(s[i]) == -1 {
			i++
			if i >= len(s) {
				return result
			}
		}
		result = append(result, toNum(s[i]))
		i++
		if i >= len(s) {
			return result
		}
		for toNum(s[i]) != -1 {
			result[len(result)-1] = result[len(result)-1]*10 + toNum(s[i])
			i++
			if i >= len(s) {
				return result
			}
		}
	}

	return result
}
