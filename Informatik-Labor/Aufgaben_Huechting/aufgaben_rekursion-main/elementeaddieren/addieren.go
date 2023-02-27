package elementeaddieren

// Liefert die Summe aller Elemente in list.
func AddElements(list []int) int {
	// TODO

	//Beobachtung: Die Summe aller elemente ist:
	// - 0 falls die liste leer ist  
	if len(list) == 0 {
		return 0
	}
	// - Das erste element 
	// - Plus die Summe aller restlichen elemente 

	head, tail := list[0],list[1:]

	

	return head + AddElements(tail)
}

func RemoveOneElement(s int,list []int) []int{

	return append(list[:s], list[s+1:]...)

}

// Liefert die Summe aller Elemente in list.
func AddElementsfor(list []int) int {
	// TODO

	result := 0 

	for _,e := range list {
	
		result += e

	}

	return result
}