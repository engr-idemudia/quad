package piscine

import (
	"github.com/01-edu/z01"
)

func QuadE(x, y int) {
	if x > 0 && y > 0 {
		z01.PrintRune('A')

		if y >= 1 {
			for a := 0; a < (x - 2); a++ {
				z01.PrintRune('B')
			}
			if x >= 2 {
				z01.PrintRune('C')
			}
			z01.PrintRune('\n')

			for b := 0; b < (y - 2); b++ {
				z01.PrintRune('B')
				for a := 0; a < (x - 2); a++ {
					z01.PrintRune(' ')
				}
				if x >= 2 {
					z01.PrintRune('B')
				}
				z01.PrintRune('\n')
			}

			if y >= 2 {
				z01.PrintRune('C')
				for a := 0; a < (x - 2); a++ {
					z01.PrintRune('B')
				}
				if x >= 2 {
					z01.PrintRune('A')
				}
				z01.PrintRune('\n')
			}
		}
	}
}
