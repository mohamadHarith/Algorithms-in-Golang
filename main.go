package main

import (
	arrays "algorithms-in-golang/Arrays"
	"log"
)

func main() {
	log.Println(arrays.TrappingRainWater1([]int{3, 4, 0, 1, 2, 4}, 6))
	log.Println(arrays.TrappingRainWater2([]int{3, 4, 0, 1, 2, 4}, 6))
}
