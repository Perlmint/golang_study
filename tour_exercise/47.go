package main

import "tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dx)
	for x, _ := range ret {
		ret[x] = make([]uint8, dy)
		for y, _ := range ret[x]{
			if x == 2 * y {
				ret[x][y] = uint8(255)
			} else {
				ret[x][y] = uint8(0)
			}
		}
	}
	
	return ret
}

func main() {
	pic.Show(Pic)
}

