package power

import "fmt"

func ExamplePower() {
	fmt.Println(Power(2, 0))
	fmt.Println(Power(2, 1))
	fmt.Println(Power(2, 2))
	fmt.Println(Power(2, 3))
	fmt.Println(Power(2, 4))
	fmt.Println(Power(2, 5))
	fmt.Println(Power(2, 14))
	fmt.Println(Power(3, 0))
	fmt.Println(Power(3, 1))
	fmt.Println(Power(3, 2))
	fmt.Println(Power(0, 1))
	fmt.Println(Power(0, 2))
	fmt.Println(Power(-1, 2))
	fmt.Println(Power(-1, 3))

	// Output:
	// 1
	// 2
	// 4
	// 8
	// 16
	// 32
	// 16384
	// 1
	// 3
	// 9
	// 0
	// 0
	// 1
	// -1
}
