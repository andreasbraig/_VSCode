package aufgabe6

/*
AUFGABENSTELLUNG:
Hier ist ein Struct Matrix vorgegeben.
Vervollständigen Sie die Funktion Mult.
PUNKTE: 15
BEWERTUNG:
*/

type Matrix struct {
	Rows [][]int
}

// Mult erwartet eine weitere Matrix m2, berechnet die Matrixmultiplikation der beiden und liefert diese zurück.
// Gehen Sie davon aus, dass die Matrizen wohlgeformt sind und von den Dimensionen her passen,
// so dass eine Multiplikation möglich ist.
func (m1 Matrix) Mult(m2 Matrix) Matrix {
	result := Matrix{Rows: [][]int{}}
	// TODO
	return result
}
