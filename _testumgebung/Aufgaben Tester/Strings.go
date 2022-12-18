package main

import "fmt"

func main() {
	s1 := "A"
	s2 := "BBB"
	//Expected ABBB

	result := ""
	// TODO
	
	var max int 

	if len(s1) > len(s2) {
		max = len(s1)
	}else {
		max = len(s2)
	}

	for i := 0 ;i < max; i++ {

		if i > len(s1)-1 {
			result = result + string(s2[i])
		}else if i > len(s2)-1 {
			result = result + string(s1[i])
		}else{
			result = result + string(s1[i]) + string(s2[i])
		}
		
	}

	fmt.Println(result)

}
