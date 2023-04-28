package aufgabe5

import "fmt"

func ExampleArrayAvg() {
	fmt.Println(ArrayAvg([]float64{1, 3, 5, 7}))
	fmt.Println(ArrayAvg([]float64{1, 1, 203, 80}))
	fmt.Println(ArrayAvg([]float64{7, 3, 1, 2}))
	fmt.Println(ArrayAvg([]float64{3, 3, 0, 2}))
	fmt.Println(ArrayAvg([]float64{}))

	// Output:
	// 4
	// 71.25
	// 3.25
	// 2
	// 0
}
