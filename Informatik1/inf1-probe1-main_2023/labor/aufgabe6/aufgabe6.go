package aufgabe6

/*
AUFGABENSTELLUNG:
Hier ist ein Struct Connection vorgegeben.
Vervollständigen Sie die zugehörige Methode IsBefore.
PUNKTE: 15
BEWERTUNG:
*/

// Connection repräsentiert eine Punkt-zu-Punkt-Verbindung
// zwischen Städten in einem Bahn-Fahrplan.
type Connection struct {
	Origin, Destination            string
	DepartureHour, DepartureMinute int
	ArrivalHour, ArrivalMinute     int
}

// IsBefore ist eine Methode von Connection und erwartet eine weitere Connection next.
// Sie prüft, ob die gegebene Verbindung conn zeitlich vollständig vor next liegt.
func (conn Connection) IsBefore(next Connection) bool {
	// TODO
	return false
}
