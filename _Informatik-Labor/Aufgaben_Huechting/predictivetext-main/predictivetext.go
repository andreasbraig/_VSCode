package main

// Erwartet eine Liste von Strings dataBase und einen einzelnen String prefix.
// Liefert eine gefilterte mit genau den Strings aus dataBase, von denen prefix
// tatsächlich ein Präfix ist. D.h. die Strings, die mit prefix beginnen.
func FilterByPrefix(dataBase []string, prefix string) []string {
	result := []string{}
	// TODO
	return result
}

// Erwartet zwei Strings s1 und s2. Prüft, ob s2 aus s1 entsteht,
// indem man ein Zeichen aus s1 weglässt.
func HasExtraLetter(s1, s2 string) bool {
	// TODO
	return false
}

// Erwartet eine Liste von Strings dataBase und einen einzelnen String substring.
// Liefert eine gefilterte mit genau den Strings aus dataBase, aus denen substring
// entsteht indem man ein Zeichen entfernt.
func FilterByExtraLetter(dataBase []string, substring string) []string {
	result := []string{}
	// TODO
	return result
}

// Erwartet zwei Strings s1 und s2. Prüft, ob s2 aus s1 entsteht,
// indem man beliebig viele Zeichen aus s1 weglässt.
func ContainsScatteredSubstring(s1, s2 string) bool {
	// TODO
	return false
}

// Erwartet eine Liste von Strings dataBase und einen einzelnen String substring.
// Liefert eine gefilterte mit genau den Strings aus dataBase, aus denen substring
// entsteht indem man beliebig viele Zeichen entfernt.
func FilterByScatteredSubstring(dataBase []string, substring string) []string {
	result := []string{}
	// TODO
	return result
}

// Hilfsfunktion: Liefert den String, der aus s entsteht, wenn man die Zeichen
// an den Stellen i und j vertauscht.
func SwitchLetters(s string, i, j int) string {
	result := ""
	// TODO
	return result
}

// Erwartet zwei Strings s1 und s2. Prüft, ob s2 aus s1 entsteht,
// indem man zwei benachbarte unterschiedliche Zeichen vertauscht.
func ContainsSwitchedLetters(s1, s2 string) bool {
	// TODO
	return false
}

// Erwartet eine Liste von Strings dataBase und einen einzelnen String substring.
// Liefert eine gefilterte mit genau den Strings aus dataBase, aus denen substring
// durch Vertauschen benachbarter Buchstaben hervorgeht
func FilterBySwitchedLetters(dataBase []string, substring string) []string {
	result := []string{}
	// TODO
	return result
}

// Hilfsfunktion: Prüft, ob der gegebene String in der Liste enthalten ist.
func ListContains(list []string, s string) bool {
	// TODO
	return false
}

// Erwartet eine Liste von Strings und liefert die gleiche Liste ohne Duplikate.
func RemoveDuplicates(list []string) []string {
	result := []string{}
	// TODO
	return result
}

// Erwartet eine Liste von Strings dataBase und einen einzelnen String incomplete.
// Liefert eine Liste von Vorschlägen zu Vervollständigung von incomplete.
// Die Vorschläge sollen in folgender Reihenfolge in der Liste vorkommen:
//
// 1. Alle gültigen Fortsetzungen von incomplete
// 2. Alle Strings, aus denen incomplete durch Buchstabedreher entstanden sein könnte
// 3. Alle Strings, bei deren incomplete eines oder mehrere Zeichen fehlen.
//
// Die Liste der Vorschläge enthält keine Duplikate!
func SuggestCompletions(dataBase []string, incomplete string) []string {
	// TODO
	return []string{}
}
