package main

import "fmt"

var AX = []int{11, 13, 12, 15, 10, -1, 14, -8, -7, -8, 11, -2, -2, -13}
var AY = []int{5, 5, 1, 15, 2, 2, 5, 8, 14, 12, 7, 14, 13, 6}
var DZ = []int{1, 1, 1, 1, 1, 26, 1, 26, 26, 26, 1, 26, 26, 26}
var maxZ []int

func main() {
	for i, _ := range DZ {
		mz := 1
		for j := i; j < len(DZ); j++ {
			mz *= DZ[j]
		}

		maxZ = append(maxZ, mz)
	}

	fmt.Printf("%v\n", simulate([]int{}, 0))

}

func simulate(path []int, z int) []int {
	i := len(path)

	for w := 1; w <= 9; w++ {
		z := step(i, w, z)

		if z > maxZ[i] {
			continue
		}

		path := append([]int{}, path...)
		path = append(path, w)

		if len(path) == 14 {
			if z == 0 {
				return path
			}
		} else {
			s := simulate(path, z)
			if s != nil {
				return s
			}
		}
	}

	return nil
}

func step(i, w, z int) int {
	ax := AX[i]
	ay := AY[i]
	dz := DZ[i]

	x := z%26 + ax

	if dz > 1 {
		z /= dz
	}

	if x != w {
		z *= 26
		z += w + ay
	}

	return z
}
