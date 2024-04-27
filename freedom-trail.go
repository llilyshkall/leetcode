package main

import (
	"fmt"
	"math"
)

func findRotateSteps(ring string, key string) int {
	ringLen := len(ring)
	keyLen := len(key)
	bestSteps := make([][]int, ringLen)
	for i := range bestSteps {
		bestSteps[i] = make([]int, keyLen+1)
		for j := range bestSteps[i] {
			bestSteps[i][j] = math.MaxInt
		}
		bestSteps[i][keyLen] = 0
	}
	for keyIndex := keyLen - 1; keyIndex >= 0; keyIndex-- {
		for ringIndex := 0; ringIndex < ringLen; ringIndex++ {
			for charIndex := 0; charIndex < ringLen; charIndex++ {
				if ring[charIndex] == key[keyIndex] {
					bestSteps[ringIndex][keyIndex] = min(bestSteps[ringIndex][keyIndex],
						1+countSteps(ringIndex, charIndex, ringLen)+bestSteps[charIndex][keyIndex+1])
				}
			}
		}
	}

	return bestSteps[0][0]
}

func countSteps(curr int, next int, ringLength int) int {
	stepsBetween := curr - next
	if stepsBetween < 0 {
		stepsBetween *= -1
	}
	stepsAround := ringLength - stepsBetween
	return min(stepsBetween, stepsAround)
}

func main() {
	fmt.Println(findRotateSteps("godding", "gd"))
	fmt.Println(findRotateSteps("godding", "godding"))
	fmt.Println(findRotateSteps("abcde", "ade"))
	fmt.Println(findRotateSteps("nyngl", "yyynnnnnnlllggg"))
}
