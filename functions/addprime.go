package functions

// "fmt"
// "os"
// "strconv"

func NewPrime(n int) bool {
	if n < 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func AddPrime(n int) int {
	result := 0
	for i := 2; i <= n; i++ {
		if NewPrime(i) {
			result += i
		}
	}
	return result
}

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println(0)
// 	}

// 	n, err := strconv.Atoi(os.Args[1])
// 	if err != nil || n < 0 {
// 		fmt.Println(0)
// 		return
// 	}
// }
