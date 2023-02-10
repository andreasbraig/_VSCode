package main

import "fmt"

func Ack(m, n int) int {


	//Ack(0,n) = n +1
	//Ack(m,0) = Ack(m-1,1)
	//Ack(m,n) = Ack(m-1,Ack(m,n-1))

	if m == 0 {
		return n+1
	}
	if n == 0 {
		return Ack(m-1,1)
	}
	return Ack(m-1,Ack(m,n-1))
}

func main() {
	fmt.Println(Ack(4, 4))
}
