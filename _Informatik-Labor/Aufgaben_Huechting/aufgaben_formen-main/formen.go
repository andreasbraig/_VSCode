package main

import "fmt"

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Das Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
func PrintSquare(n int) {
	// TODO
	for i := 1; i <= n; i++ { //Anzahl der Zeilen
		for j := 0; j < n; j++ { //Anzahl der Sterne Pro Zeile
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Der Rand des Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
// Das Innere soll aus Leerzeichen bestehen.
func PrintEmptySquare(n int) {
	// TODO
	for i := 0; i < n; i++ { // erste Zeile
		fmt.Print("*")
	}
	fmt.Println()               // umbruch
	for i := 1; i <= n-2; i++ { //Anzahl der Zeilen die hohl sind
		fmt.Print("*")
		for j := 0; j < n-2; j++ { //Sterne Pro Zeile
			fmt.Print(" ")
		}
		fmt.Print("*")
		fmt.Println() //umbruch nach Zeile
	}
	for i := 0; i < n; i++ { // erste Zeile
		fmt.Print("*")
	}
	fmt.Println()
}

// Erwartet eine Zahl  n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Die Zeichen für Rand und Inneres werden als Parameter angegeben.
func PrintCustomSquare(n int, border, fill string) {
	// TODO
	for i := 0; i < n; i++ { // erste Zeile
		fmt.Print(border)
	}
	fmt.Println()               // umbruch
	for i := 1; i <= n-2; i++ { //Anzahl der Zeilen die hohl sind
		fmt.Print(border)
		for j := 0; j < n-2; j++ { //Sterne Pro Zeile
			fmt.Print(fill)
		}
		fmt.Print(border)
		fmt.Println() //umbruch nach Zeile
	}
	for i := 0; i < n; i++ { // erste Zeile
		fmt.Print(border)
	}
	fmt.Println()
}

// Erwartet zwei Zahlen m und n und gibt ein Rechteck
// der Breite m und der Höhe n auf die Konsole aus.
// Das Rechteck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintRectangle(m, n int) {
	// TODO
	for i := 1; i <= n; i++ { //Anzahl der Zeilen
		for j := 0; j < m; j++ { //Anzahl der Sterne Pro Zeile
			fmt.Print("*")
		}
		fmt.Println()
	}

}

// Erwartet eine Zahl n und gibt ein Dreieck auf die Konsole aus.
// Das Dreieck soll ein halbes n x n-Quadrat sein, das auf der
// linken oberen Seite ausgefüllt ist (siehe Test).
// Das Dreieck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintTriangle(n int) {
	// TODO
	for i := 1; i <= n; i++ { //Anzahl der Zeilen
		for j := 0; j < n; { //Anzahl der Sterne Pro Zeile
			fmt.Print("*")
		}
		fmt.Println()
		n = n - 1
	}
}
