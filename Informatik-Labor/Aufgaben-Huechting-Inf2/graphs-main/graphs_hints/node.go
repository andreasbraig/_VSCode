package graphs

// Node is a struct containing a label and a slice of edges.
// It is used to represent a node in a graph.
type Node struct {
	Label      string
	neighbours []*Node
}

// New creates a new Node.
func NewNode(label string) *Node {
	return &Node{Label: label}
}

// String returns the label of the node.
func (n *Node) String() string {
	return n.Label
}

// NeighbourCount returns the number of neighbours of a node.
func (n *Node) NeighbourCount() int {
	return len(n.neighbours)
}

// GetNeighbour returns the neighbour with the given label.
// If no neighbour with the given label exists, nil is returned.
func (n *Node) GetNeighbour(label string) *Node {
	// TODO
	return nil

	// Laufe durch alle Nachbarn und prüfe, ob der Knoten das gesuchte Label hat.
}

// AddNeighbour adds a neighbour with a label to a node.
// If a Neighbour with the same label already exists, it is not added.
func (n *Node) AddNeighbour(label string) {
	// TODO

	// Erstelle einen neuen Knoten mit dem Label und füge ihn als Nachbarn hinzu.
	// Nutze dafür die Methode AddNeighbourNode.
}

// AddNeighbourNode adds node m as a neighbour.
// If a node with that label already exists, it is not added.
// If m is nil, it is not added.
func (n *Node) AddNeighbourNode(m *Node) {
	// TODO

	// Füge m zu n.neighbours hinzu.
	// Prüfe vorher, ob der neue Knoten überhaupt existiert und ob noch kein
	// Knoten mit diesem Label vorhanden ist.
	// Die letztere Prüfung kann mittels GetNeighbour erfolgen.
}

// NeighboursMaxDistance returns all neighbours of a node with a distance of at most maxDistance.
func (n *Node) NeighboursMaxDistance(maxDistance int) []*Node {
	result := []*Node{}
	// TODO

	return result

	// Erstelle eine Liste für die Nachbarn und füge alle direkten Nachbarn (also Zielknoten der Kanten) hinzu.
	// Falls maxDistance > 1 ist, rufe diese Funktion rekursiv mit jedem der direkten Nachbarn und maxDistance-1 auf.
	// Für die Zusatzanforderung aus den Tests mit Graphen, die Kreise enthalten: Entferne Duplikate aus der Liste.
}

// CanReachLabel_MaxDepth expects a label and a maximal search depth.
// returns true if a node with that label is reachable.
func (n *Node) CanReachLabel_MaxDepth(label string, maxDepth int) bool {
	// TODO
	return false

	// Rekursive Tiefensuche
	// Prüfe, ob der aktuelle Knoten das gesuchte Label hat oder ob maxDepth == 0 ist.
	// Dies sind die Basisfälle, in denen die Suche abgebrochen wird.
	// Falls nicht, rufe diese Funktion rekursiv mit jedem der direkten Nachbarn und maxDepth-1 auf.
}

// CanReachLabel expects a label.
// returns true if a node with that label is reachable in any number of steps.
func (n *Node) CanReachLabel(label string) bool {
	// TODO
	return false

	// Im Gegensatz zu CanReachLabel_MaxDepth wird hier nicht rekursiv vorgegangen.
	// Für die Knoten, die tatsächlich erreichbar sind, würde ein rekursiver Ansatz
	// ähnlich wie bei CanReachLabel_MaxDepth funktionieren.
	// Für Knoten, die nicht erreichbar sind, bricht dann aber die Rekursion nicht ab.

	// Daher wird hier ein iterativer Ansatz gewählt:
	// Erstelle zwei Listen:
	// * lastneighbours enthält alle Knoten, die in der letzten Iteration gefunden wurden.
	// * currentneighbours enthält alle Knoten, die in der aktuellen Iteration gefunden wurden.
	// Die Suche berechnet nun immer weiter die Nachbarn der Nachbarn, bis sich die beiden Listen nicht mehr unterscheiden.
	// Dazu kann die Funktion NeighboursMaxDistance in einer Schleife verwendet werden.
	// Wenn die beiden Listen gleich sind, ist die Suche abgeschlossen,
	// da dann keine neuen Knoten mehr gefunden werden können.
	// Prüfe nun, ob der gesuchte Knoten in der Liste der aktuellen Nachbarn enthalten ist.
}
