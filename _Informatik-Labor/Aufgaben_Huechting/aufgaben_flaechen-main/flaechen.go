package main

import "math"

// Erwartet zwei Seitenlängen a und b.
// Liefert die Fläche des entsprechenden Rechtecks.
func AreaRectangle(a, b float64) float64 {
	// TODO
	return a * b //übergabe und Rechnung des Wertes
}

// Erwartet eine Seitenlänge a.
// Liefert die Fläche des entsprechenden Quadrats.
func AreaSquare(a float64) float64 {
	// TODO
	return a * a //Übergabe und Rechnung des Wertes
}

// Erwartet zwei Seitenlängen a und b.
// Liefert die Fläche eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func AreaRightTriangle(a, b float64) float64 {
	// TODO

	return a * b * 0.5
}

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang des entsprechenden Rechtecks.
func PerimeterRectangle(a, b float64) float64 {
	// TODO
	return 2*a + 2*b
}

// Erwartet eine Seitenlänge a.
// Liefert den Umfang des entsprechenden Quadrats.
func PerimeterSquare(a float64) float64 {
	// TODO
	return 4 * a
}

// Erwartet die Längen der Katheten eines rechtwinkligen Dreiecks.
// Liefert die Länge der Hypotenuse.
func Hypotenuse(a, b float64) float64 {
	// Hinweis: Hierfür ist die Funktion
	//          math.Sqrt() nützlich.
	// TODO
	return math.Sqrt(a*a + b*b)

}

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func PerimeterRightTriangle(a, b float64) float64 {
	// TODO
	return a + b + Hypotenuse(a, b)
}
