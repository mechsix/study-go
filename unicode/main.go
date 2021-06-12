package main

import "fmt"

func main() {
	str := "12345測試"

	fmt.Printf("len %d: %s\n", len(str), str)
	fmt.Printf("Characters:\n")

	for index, unicodeChar := range str {
		fmt.Printf("(%d) %%s: %s, %%c: %c\n", index, unicodeChar, unicodeChar)
	}
}
