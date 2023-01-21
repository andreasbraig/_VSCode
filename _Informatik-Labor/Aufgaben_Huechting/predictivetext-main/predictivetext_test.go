package main

import "fmt"

func ExampleFilterByPrefix() {
	db1 := []string{
		"Hallo", "Haus", "Wellenreiten", "Hamburg", "Westpfalz", "Kohle", "Hallenbad"}

	fmt.Println(FilterByPrefix(db1, "Halle"))
	fmt.Println(FilterByPrefix(db1, "Hall"))
	fmt.Println(FilterByPrefix(db1, "Hallo"))
	fmt.Println(FilterByPrefix(db1, "Welle"))
	fmt.Println(FilterByPrefix(db1, "West"))
	fmt.Println(FilterByPrefix(db1, "Kohl"))

	// Output:
	// [Hallenbad]
	// [Hallo Hallenbad]
	// [Hallo]
	// [Wellenreiten]
	// [Westpfalz]
	// [Kohle]
}

func ExampleHasExtraLetter() {
	fmt.Println(HasExtraLetter("Fehler", "Feler"))
	fmt.Println(HasExtraLetter("Fehler", "Fehle"))
	fmt.Println(HasExtraLetter("A", ""))
	fmt.Println(HasExtraLetter("Fehler", "Felerr"))
	fmt.Println(HasExtraLetter("B", "C"))

	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleFilterByExtraLetter() {
	db1 := []string{"Haus", "Hals", "Taub", "Taube", "Kosten", "Knoten"}

	fmt.Println(FilterByExtraLetter(db1, "Has"))
	fmt.Println(FilterByExtraLetter(db1, "Tau"))
	fmt.Println(FilterByExtraLetter(db1, "Koten"))

	// Output:
	// [Haus Hals]
	// [Taub]
	// [Kosten Knoten]
}

func ExampleContainsScatteredSubstring() {
	fmt.Println(ContainsScatteredSubstring("Fehler", "Feler"))
	fmt.Println(ContainsScatteredSubstring("Fehler", "Felr"))
	fmt.Println(ContainsScatteredSubstring("Fehler", "Fhr"))
	fmt.Println(ContainsScatteredSubstring("Fehler", "Fehler"))
	fmt.Println(ContainsScatteredSubstring("Fehler", "Feeehhler"))

	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleFilterByScatteredSubstring() {
	db1 := []string{"Haus", "Hals", "Taub", "Taube", "Kosten", "Knoten"}

	fmt.Println(FilterByScatteredSubstring(db1, "Has"))
	fmt.Println(FilterByScatteredSubstring(db1, "Hs"))
	fmt.Println(FilterByScatteredSubstring(db1, "Ta"))
	fmt.Println(FilterByScatteredSubstring(db1, ""))

	// Output:
	// [Haus Hals]
	// [Haus Hals]
	// [Taub Taube]
	// [Haus Hals Taub Taube Kosten Knoten]
}

func ExampleSwitchLetters() {
	fmt.Println(SwitchLetters("AB", 0, 1))
	fmt.Println(SwitchLetters("ABCDE", 2, 3))
	fmt.Println(SwitchLetters("ABCDE", 1, 3))

	// Output:
	// BA
	// ABDCE
	// ADCBE
}

func ExampleContainsSwitchedLetters() {
	fmt.Println(ContainsSwitchedLetters("ABC", "ACB"))
	fmt.Println(ContainsSwitchedLetters("ACB", "ABC"))
	fmt.Println(ContainsSwitchedLetters("ABC", "ABC"))
	fmt.Println(ContainsSwitchedLetters("AA", "AA"))
	fmt.Println(ContainsSwitchedLetters("ABCDE", "ABC"))
	fmt.Println(ContainsSwitchedLetters("ABC", "ABCDE"))

	// Output:
	// true
	// true
	// false
	// false
	// false
	// false
}

func ExampleFilterBySwitchedLetters() {
	db1 := []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}

	fmt.Println(FilterBySwitchedLetters(db1, "ABC"))
	fmt.Println(FilterBySwitchedLetters(db1, "ACB"))
	fmt.Println(FilterBySwitchedLetters(db1, "BAC"))
	fmt.Println(FilterBySwitchedLetters(db1, "BCA"))
	fmt.Println(FilterBySwitchedLetters(db1, "CAB"))
	fmt.Println(FilterBySwitchedLetters(db1, "CBA"))

	// Output:
	// [ACB BAC]
	// [ABC CAB]
	// [ABC BCA]
	// [BAC CBA]
	// [ACB CBA]
	// [BCA CAB]
}

func ExampleListContains() {
	db1 := []string{"Haus", "Hals", "Halstuch", "Handtuch", "Staubtuch"}

	fmt.Println(ListContains(db1, "Haus"))
	fmt.Println(ListContains(db1, "Halstuch"))
	fmt.Println(ListContains(db1, "Küchentuch"))

	// Output:
	// true
	// true
	// false
}

func ExampleRemoveDuplicates() {
	db1 := []string{"Haus", "Auto", "Haus", "Hand", "Auto"}

	fmt.Println(RemoveDuplicates(db1))

	// Output:
	// [Haus Auto Hand]
}

func ExampleSuggestCompletions() {
	db1 := []string{"Haus", "Hals", "Halstuch", "Handtuch", "Staubtuch"}

	fmt.Println(SuggestCompletions(db1, "Hal"))
	fmt.Println(SuggestCompletions(db1, "Ha"))
	fmt.Println(SuggestCompletions(db1, "tuch"))
	fmt.Println(SuggestCompletions(db1, "taub"))

	// Output:
	// [Hals Halstuch]
	// [Haus Hals Halstuch Handtuch]
	// [Halstuch Handtuch Staubtuch]
	// [Staubtuch]
}
