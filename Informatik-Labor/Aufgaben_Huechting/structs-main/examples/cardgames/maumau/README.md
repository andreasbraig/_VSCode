# Package `maumau`

In diesem Package werden Datentypen und Funktionen für ein Mau-Mau-Spiel definiert.

Es gibt einen Datentyp [`Game`](game.go), der die notwendigen Daten für ein
Mau-Mau-Spiel enthält.
Dies sind ein Kartenstapel, ein Ablagestapel und eine Liste von Spielern.
Diese werden in einem Struct kombiniert.

Der `Game`-Typ stellt einige Methoden bereit, die das Spiel steuern.
Es gibt Methoden `Setup` und `Run` für Aufbau und Ablauf des Spiels,
sowie Methoden, mit denen der Zustand abgefragt und ausgegeben werden kann.

Im Unterverzeichnis [`main`](main/maumau.go) gibt es außerdem ein Package `main`
mit einem ausführbaren Programm, das das tatsächliche Spiel startet.

## Anmerkung und weitere Aufgaben

Dieses Spiel ist noch nicht vollständig.
Das [`main`](main/maumau.go)-Programm legt bisher hartcodiert zwei Spieler an,
gibt ihnen Karten und fragt sie dann abwechselnd nach ihren Zügen und führt diese aus.
Es prüft bisher nicht, ob die Spieler gültige Eingaben machen und wird ggf. abstürzen,
falls sie es nicht tun. Es prüft auch nicht, ob die Züge nach den Mau-Mau-Regeln erlaubt
sind. Es gibt also eine Reihe weiterer Aufgaben zu erledigen:

* Überprüfung der Eingabe bei der Abfrage der Spielzüge
* Überprüfung, ob ein Spielzug erlaubt ist
* Karte Nachziehen, falls ein Spieler nicht legen kann
* Neu mischen, falls der Kartenstapel leer ist
* Fehlerbehandlung, falls alle Karten an Spieler verteilt sind und gezogen werden muss
* Umsetzung der Sonderregeln bei Mau-Mau: Bei 7 zwei ziehen, bei 8 aussetzen etc.
