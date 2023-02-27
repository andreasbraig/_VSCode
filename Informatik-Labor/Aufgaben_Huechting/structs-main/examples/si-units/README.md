# Si-Einheitensystem

Hier wird ein einfaches Einheitensystem definiert.

* [`Duration`](duration/duration.go) beschreibt Zeiträume.
  Eine `Duration` kann auf verschiedene Weisen aus
  Sekunden-, Minuten-, Tages- etc. Angaben erzeugt werden.
* [`Distance`](distance/distance.go) beschreibt Entfernungen.
  Eine `Distance` kann auf verschiedene Weisen aus
  Meter-, Kilometer-, Meilen- etc. Angaben erzeugt werden.
* [`Velocity`](veloctiy/velocity.go) beschreibt Geschwindigkeiten.
  Eine `Velocity` kann entweder aus `Duration`- und `Distance`-Objekten
  oder aus Zahlenangaben gebildet werden.

`Duration` und `Distance`
sind dabei einfache Typdeklarationen, die den jeweiligen Datentyp direkt auf `int`
abbilden. `Velocity` ist dann ein Struct, das aus
`Duration` und `Distance` zusammengesetzt ist.
