package hanoi

import "fmt"

// Problemlösung des Rätsels der Türme von Hanoi

func BewegePlatte(start,Ziel string) {

	fmt.Printf("%s ==> %s",start,Ziel)

}


//Bewegt einen Turm der höhe 1
func Hanoi1(start,mitte,ziel string){
	BewegePlatte(start,ziel)
}

//Bewegt einen Turm der Höhe 2
func Hanoi2(start,mitte,ziel string) {
	Hanoi1(start,ziel,mitte)
	BewegePlatte(start,ziel)
	Hanoi1(mitte,start,ziel)
}

//Bewegt einen turm der Höhe 3
func Hanoi3(start,mitte,ziel string) {
	Hanoi2(start,ziel,mitte)
	BewegePlatte(start,ziel)
	Hanoi2(mitte,start,ziel)
}

//Rekursive Problemlösung
//Man muss nur die Höhe der türme h übergeben 
func Hanoi(h int, start,mitte,ziel string) {
	if h == 0 {
		return 
	}
	Hanoi(h-1,start,ziel,mitte)
	BewegePlatte(start,ziel)
	Hanoi(h-1,mitte,start,ziel)
}