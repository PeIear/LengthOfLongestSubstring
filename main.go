package main

import (
	"fmt"
)

func main() {

	fmt.Println(LengthOfLongestSubstring("abcabcabc"))

}

func LengthOfLongestSubstring(s string) int {
	result := 0
	type Val struct {
		Str string
		Len int
	}
	NewSlice := make([]Val, 0)
	for i := 0; i <= len(s)-1; i++ {
		str := ""
		j := strings.IndexRune(s[i+1:], rune(s[i]))
		if j > 0 {
			str = s[i : j+i+1]
		} else {
			str = string(s[i])
		}
		NewSlice = append(NewSlice, Val{str, len(str)})
	}
	sort.SliceStable(NewSlice, func(i, j int) bool {
		return NewSlice[i].Len > NewSlice[j].Len
	})
	result = NewSlice[0].Len
	fmt.Println(NewSlice[0].Str)
	return result
}
