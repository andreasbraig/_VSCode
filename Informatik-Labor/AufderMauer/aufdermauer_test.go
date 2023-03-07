package aufdermauer

import "fmt"

func ExampleTanzeWanze() {

	TanzeWanze()

	//Output:

}

func ExampleKillWord() {

	KillWord("Wanze")
	fmt.Println()
	BuildWord("", "Wanze")

	fmt.Println()

	
	KillWord("tanzen")
	fmt.Println()
	BuildWord("", "tanzen")
	//Output:

}

func ExampleAufderMauer() {

	AufderMauer()


	//Output:
}
