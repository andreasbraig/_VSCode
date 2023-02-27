package matrizen

import "fmt"

func ExampleMatrizen() {

	//Erzeuge eine 2D-Liste aus Zahlen

	m1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // Ist eine Liste aus listen

	fmt.Println(m1)

	fmt.Println(m1[0])
	fmt.Println(m1[0][0])
	fmt.Println(m1[2][0])

	for _, row := range m1{
		fmt.Println(row)
	}
	for _, row := range m1{
		for pos,element := range row {
			fmt.Print(element)
			if pos < len(row) -1 {
				fmt.Print(",")
			}
		}
		fmt.Println()
	}

	//Output:


}
