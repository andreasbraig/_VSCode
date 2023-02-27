package gameboard

import "fmt"

type Row []string

// Liefert true falls Row nur aus symbol besteht
func (row Row) ContainsOnly(symbol string) bool {

	for _, character := range row {
		if character != symbol {
			return false
		}
	}
	return true

}

// Liefert eine String-Representation der Zeile
func (row Row) String() string {

	result := "|"

	for _, element := range row {
		result += fmt.Sprintf(" %s |", element)
	}

	return result

}
