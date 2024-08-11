package utils

func ToNum(c byte) int {
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
		for ToNum(s[i]) == -1 {
			i++
			if i >= len(s) {
				return result
			}
		}
		result = append(result, ToNum(s[i]))
		i++
		if i >= len(s) {
			return result
		}
		for ToNum(s[i]) != -1 {
			result[len(result)-1] = result[len(result)-1]*10 + ToNum(s[i])
			i++
			if i >= len(s) {
				return result
			}
		}
	}

	return result
}
