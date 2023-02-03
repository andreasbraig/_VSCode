package gameboard

type Board []Row


// Liefert true, falls irgendeine Zeile in board mit
// symbol gefüllt ist.
func (board Board)AnyRowContainsOnly(symbol string) bool { //In der Funktion gibt er jetzt den Datentyp an. (Methode)

	for _, row := range board {

		if row.ContainsOnly(symbol) { // Ruft die Funktion  Contains only im file row aus 
			return true
		}
	}
	return false
}
