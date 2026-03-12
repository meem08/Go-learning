package functions

func IsPrime(n int) bool {
	if n < 1 {
		return 0
	}

	for i :=2; i*i <= n; i++{
		if n%i == 0 {
			return false
		}
	}
	return true
}

func AddPrime(n int)int {
	result := 0
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			result += i
		}
	}
}

func main () {
	if len(os.Args) != 2 {
		fmt.Println(0)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Println(0)
		return
	} 
}

