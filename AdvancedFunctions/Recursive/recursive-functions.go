package recursive

func Fibonacci(position uint) uint {
	// fmt.Println("Funções Recursivas")
	if position <= 1 {
		return position
	}

	return Fibonacci(position-2) + Fibonacci(position-1)
}
