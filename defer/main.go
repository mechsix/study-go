package main
import (
	"fmt"
)

func level1 (i int) int {
	fmt.Printf("Level 1 enter: %d\n", i)
	// Defer called "after" func Level1 return
	defer level2(i)
	fmt.Printf("Level 1 exit: %d\n", i)
	return i
}

func level2 (i int) {
	fmt.Printf("Level 2 enter: %d\n", i)
	fmt.Printf("Level 2 exit: %d\n", i)
}

func main() {
	fmt.Printf("Returned %d\n", level1(42))
}
