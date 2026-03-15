// Write the functions that prints in ascending order on a single all unit combination of three different digits so that the first digit
// is lower than the second, and the second is lower than the third. These combinations are separated by comma and a space

package functions

import "fmt"

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				fmt.Print(i)
				fmt.Print(j)
				fmt.Print(k)
				if i != 7 || j != 8 || k != 9 {
					fmt.Print(",")
					fmt.Print(" ")
				}
			}
		}
	}
}

