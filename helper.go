package main

func maxLen(s []string) int {
	max := len(s[0])
	for _, v := range s {
		if len(v) > max {
			max = len(v)
		}
	}
	return max
}
