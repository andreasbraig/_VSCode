package strings

// StartsWith liefert true, falls der string s mit der Sequenz seq beginnt.
func StartsWith(s, seq string) bool {
	// TODO

	// Hinweis:
	// Unterscheiden Sie drei Fälle:
	// 1. seq ist leer - Dann ist fängt s damit an.
	if seq == "" {
		return true 
	}
	// 2. s ist leer - Dann ist seq nicht der Anfang.

	if s == "" {
		return false
	}
	// 3. Beide sind nicht leer:
	//    Dann beginnt s mit seq, wenn die ersten Buchstaben gleich sind
	//    und der Rest von s mit dem Rest von seq beginnt.

	return s[0] == seq[0] && StartsWith(s[1:],seq[1:])
}
