package arztpraxis

// Medikamentenvariable
type Medikament struct {
	Name                 string
	Unvertraeglichkeiten []Medikament //Bekannte Wechselwirkungen zwischen Medikamenten

}

// AddUnverträglichkeit fügt ein anderes Medikament zur Liste der Unverträglichkeiten vion m Hinzu
func (m *Medikament) AddUnvertraeglichkeit(m2 Medikament) { //m2 ist die neue Unverträglichkeit

	m.Unvertraeglichkeiten = append(m.Unvertraeglichkeiten, m2)

}

//Prüfe ob anderes Medi (Liste der patientenmedikamente) in der Liste der Unverträglichkeiten von m vorkommt
func (m Medikament) IstVertraeglich(other Medikament) bool {

	for _,u := range m.Unvertraeglichkeiten {
		if other.Name == u.Name {
			return false
		}
	}
	return true

}