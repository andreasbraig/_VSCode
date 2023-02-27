package fibonacci

// Fib berechnet den n-ten Wert der Fibonacci-Folge.
func Fib(n int) int {
	// TODO
	if n <= 1 {
		return 1
	}
	return Fib(n-2)+Fib(n-1)
}
