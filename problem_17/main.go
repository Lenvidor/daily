package main

import (
	"fmt"
)

func main() {
	fmt.Println(Parse("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
}

func parseLine(s string) (int, int, int, bool) {
	level := 0
	wordLength := 0
	isFile := false

	for i, n := 0, len(s); i < n; i = i + 0 {
		switch s[i] {
		case '\n':
			i++
			return level, wordLength, i, isFile
		case '\t':
			level++
			i = i + 1
		default:
			if s[i] == '.' {
				isFile = true
			}
			wordLength++
			i++
		}
	}
	return level, wordLength, len(s), isFile
}

func Parse(s string) int {
	m := make(map[int]int)
	maxLength := 0
	totalLength := 0

	for i, n := 0, len(s); i < n; i = i + 0 {
		level, wordLength, newIndex, isFile := parseLine(s[i:])
		i = i + newIndex
		if level+1 <= len(m) {
			if level == 0 {
				totalLength = 0
				m = make(map[int]int)
			} else {
				for level+1 <= len(m) {
					totalLength = totalLength - m[len(m)-1] - 1
					delete(m, len(m)-1)
				}
			}
		}

		delim := 0
		if level > 0 {
			delim = 1
		}
		tempLength := totalLength + wordLength + delim
		if isFile {
			if maxLength < tempLength {
				maxLength = tempLength
			}
		} else {
			totalLength = tempLength
			m[level] = wordLength
		}

	}
	return maxLength
}
