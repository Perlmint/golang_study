package main

import (
	"tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := map[string]int{}
	for _, word := range strings.Fields(s) {
		val, exist := ret[word]
		if exist {
			ret[word] = val + 1
		} else {
			ret[word] = 1
		}
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}

