package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	// fmt.Println(test([]int{1, 2, 9}))
}

func findAnagrams(s string, p string) []int {
	res := []int{}

	if len(s) < len(p) {
		return res
	}

	pDict := make([]int, 26)
	sDict := make([]int, 26)

	for i := 0; i < len(p); i++ {
		pDict[p[i]-'a']++
		sDict[s[i]-'a']++
	}

	fmt.Println(pDict)
	fmt.Println(sDict)

	start := 0
	if isEqual(pDict, sDict) {
		res = append(res, start)
	}

	for end := len(p); end < len(s); end++ {
		fmt.Println(end, start)
		sDict[s[end]-'a']++
		sDict[s[start]-'a']--
		// fmt.Println("end", end)
		// fmt.Println("start", start)
		// fmt.Println("================================")
		// fmt.Println(sDict)
		// fmt.Println(pDict)
		start++
		if isEqual(pDict, sDict) {
			res = append(res, start)
		}
	}

	return res
}

func isEqual(pDict, sDict []int) bool {
	for i := 0; i < 26; i++ {
		if pDict[i] != sDict[i] {
			return false
		}
	}
	return true
}

func test(t []int) []int {
	r := make([]int, len(t))
	for i := len(t) - 1; i >= 0; i-- {
		if t[i] == 9 {
			r[i] = 0
			fmt.Println(r[i-1])
			r[i-1] = t[i-1] + 1
			fmt.Println(r[i-1])
			continue
		}

		if i == len(t)-1 {
			r[i] = t[i] + 1
			continue
		}

		r[i] = t[i]
	}

	return r
}
