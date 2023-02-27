package sports

// Match ist ein Datentyp, der Metadaten zu einer Sportveranstaltung speichern soll.
type Match struct {
	Home     string // Name der Heimmannschaft
	Visitors string // Name der Auswärtsmannschaft

	Location string // Name des Austragungsorts (Stadt, Stadion o.Ä.)
	Result   int    // 0: Unentschieden, 1: Heimmannschaft gewinnt, 2: Gäste gewinnen
}

// SetHomeTeam erwartet einen String und schreibt ihn in das Attribut Home.
func (m *Match) SetHome(name string) {
	m.Home = name
}
