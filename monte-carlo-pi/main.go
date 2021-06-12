package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var CYCLE = 100000

func main() {
	var hit = 0

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0 ; i < CYCLE ; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if (math.Pow(x, 2) + math.Pow(y, 2)) < 1 {
			hit++
		}
	}

	fmt.Printf("Cycle\t%d\nHit\t%d\nPi\t%f\n", CYCLE, hit, 4 * float64(hit) / float64(CYCLE) )
}
