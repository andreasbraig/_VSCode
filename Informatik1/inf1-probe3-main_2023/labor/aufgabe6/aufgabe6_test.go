package aufgabe6

import "fmt"

func ExampleMatrix_Mult_square() {
	m1 := Matrix{
		[][]int{
			{0, 1, 0},
			{0, 1, 0},
			{0, 1, 0},
		},
	}
	m2 := Matrix{
		[][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	}
	m3 := Matrix{
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}
	m4 := Matrix{
		[][]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}

	fmt.Println(m1.Mult(m2))
	fmt.Println(m1.Mult(m3))
	fmt.Println(m3.Mult(m4))

	// Output:
	// {[[1 1 1] [1 1 1] [1 1 1]]}
	// {[[4 5 6] [4 5 6] [4 5 6]]}
	// {[[1 2 3] [4 5 6] [7 8 9]]}
}

func ExampleMatrix_Mult_nonsquare() {
	m1 := Matrix{
		[][]int{
			{1, 0},
			{0, 1},
			{1, 0},
		},
	}
	m2 := Matrix{
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	}

	fmt.Println(m1.Mult(m2))

	// Output:
	// {[[1 2 3 4] [5 6 7 8] [1 2 3 4]]}
}
