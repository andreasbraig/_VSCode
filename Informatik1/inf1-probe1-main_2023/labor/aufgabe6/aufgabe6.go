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

	//Arrival von conn muss vor Departure von next Conn Liegen
	// -> next  Departure Hour muss vor conn Arrival Hour liegen

	if next.DepartureHour > conn.ArrivalHour {

		return true

	} else if next.DepartureHour == conn.ArrivalHour {
		
		// Wenn beide gleich sind wird Departure Minute mit Arrival Minute verglichen.
		
		if next.DepartureMinute > conn.ArrivalMinute {

			return true

		} else if next.DepartureMinute == next.ArrivalMinute {
			return false
		}
	}

	return false

}
