package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	// TODO

	// Hinweis:
	// Unterscheiden Sie drei Fälle:
	// 1. seq ist leer - Dann ist es immer in s enthalten.
	if seq == "" {
		return true
	}
	// 2. s ist leer - Dann ist seq auf keinen Fall enthalten.
	if s == "" {
		return false
	}
	// 3. Beide sind nicht leer:
	//    Dann ist seq enthalten, wenn s mit seq beginnt, oder der Rest von s seq enthält.

	return StartsWith(s, seq) || Contains(s[1:], seq)
}
