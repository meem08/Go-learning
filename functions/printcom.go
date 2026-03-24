package main

/*Write a function that prints, in ascending order and on a single line: all unique combinations of three different
digits so that, the first digit is lower than the second, and the second is lower than the third.
These combinations are separated by a comma and a space.*/
import "fmt"

func main() {
	PrintComb()
}

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				// if i != 0 && j != 0 && k != 0 {

				// }
				fmt.Print(i)
				fmt.Print(j)
				fmt.Print(k)

				if i != 7 || j != 8 || k != 9 {
					fmt.Print(", ")
				}
			}
		}
	}
	fmt.Println()
}
