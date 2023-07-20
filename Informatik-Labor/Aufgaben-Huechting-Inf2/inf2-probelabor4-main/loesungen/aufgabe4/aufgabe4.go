package aufgabe4

/* ERREICHBARE PUNKTE: 20 */

// In der Datei graph.go finden Sie eine Implementierung eines Knotens in einem Graphen.
// Vervollständigen Sie die Methode ReachableNodes.
// Hinweis: Sie können die Hilfsfunktion Contains verwenden.

/***HILFSFUNKTIONEN***/
func Contains(list []*Node, node *Node) bool {
	for _, n := range list {
		if n == node {
			return true
		}
	}
	return false
}

/***AUFGABE***/
// ReachableNodes soll eine Liste aller von n aus erreichbaren Knoten liefern.
func (n *Node) ReachableNodes() []*Node {
	visitedNodes := []*Node{}
	newNodes := []*Node{n}

	for len(newNodes) > 0 {
		current := newNodes[len(newNodes)-1]
		visitedNodes = append(visitedNodes, current)
		newNodes = newNodes[:len(newNodes)-1]
		for _, neighbour := range current.neighbours {
			if !Contains(visitedNodes, neighbour) {
				newNodes = append(newNodes, neighbour)
			}
		}
	}
	return visitedNodes
}
