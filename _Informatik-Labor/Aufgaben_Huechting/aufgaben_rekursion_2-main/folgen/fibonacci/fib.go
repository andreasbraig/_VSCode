package fibonacci

//Fibonacci Folge ist eine Zahl die immer Aus der Summe der beiden Zahlen davor besteht 



// Fib berechnet den n-ten Wert der Fibonacci-Folge.
func Fib(n int) int {
	// TODO

	if n <= 1 {
		return 1
	}


	return Fib(n-1)+Fib(n-2)
}
