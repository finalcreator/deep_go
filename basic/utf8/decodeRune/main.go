package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data := []byte("  hi你好ok不用谢 dd!")
	start := 0
	spaceCount := 0
	spaceLimit := 2
	for width := 0; start < len(data); start += width {
		fmt.Printf("start:=%d\n", start)
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		fmt.Printf("r=%d(%s), width=%d\n", r, string(r), width)
		if isSpace(r) {
			spaceCount++
		}
		if spaceCount > spaceLimit {
			fmt.Printf("\nreach spaceLimit = %d\n", spaceLimit)
			break
		}
	}

}

// isSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}
