package arztpraxis

// Patientenvariable
type Patient struct {
	Name  string
	Medis []Medikament //eingenommene Medikamente
	Unvis []Medikament //bekannte Unverträglichkeiten
}

// ChecMedi prüft ob der Patient p das Medikament m einnehmen darf.
func (p Patient) CheckMedi(m Medikament) bool {

	return p.PatientMediInterference(m) && p.CheckAlergies(m)

}

// MediInterference Prüft wechselwirkungen der Medikamente die ein Patient bereits nimmt gegeneinander
func (p Patient) PatientMediInterference(m Medikament) bool {

	for _, anderesMedi := range p.Medis {

		if !m.IstVertraeglich(anderesMedi) {
			return false
		}

	}
	return true

}

// CheckAlergies Überprüft die Patientenspezifischen Unverträglichkeiten und das Zugewiesene Medikament
func (p Patient) CheckAlergies(m Medikament) bool {

	for _, anderesMedi := range p.Unvis {

		if !m.IstVertraeglich(anderesMedi) {
			return false
		}

	}
	return true
}
