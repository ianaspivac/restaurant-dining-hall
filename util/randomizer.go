package util

import (
	"fmt"
	"math/rand"
)

func Randomizer() {
	var res bool
	min := 1
	max := 100
	if (rand.Intn(max-min+1)+min)%2 == 0 {
		res = false
	} else {
		res = true
	}
	fmt.Println(res)
}
func RandomizeNr(maximum int) int {
	return rand.Intn(maximum) + 1
}
