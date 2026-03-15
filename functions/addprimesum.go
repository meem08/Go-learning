package functions

// Write a program that takes a positive integer as argument and
// displays the sum of all prime numbers inferior or equal to it followed by a newline ('\n').
// If the number of arguments is different from 1, or if the argument is not
// a positive number, the program displays 0 followed by a newline

// "fmt"
// "os"
// "strconv"

// func main() {

// 	if len(os.Args) != 2 {
// 		fmt.Println(0)
// 		return
// 	}
// 	n, err := strconv.Atoi(os.Args[1])
// 	if err != nil || n < 0 {
// 		fmt.Println(0)
// 		return
// 	}
// 	fmt.Println(AddPrimeSum(n))
// }

func AddPrimeSum(n int) int {

	result := 0
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			result += i
		}
	}
	return result
}
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
