# Package `sports`

Hier wird ein Datentyp [`Match`](match.go) definiert.
Dieser Datentyp enthält Metadaten zu Sportveranstaltungen zwischen zwei Mannschaften,
wie sie z.B. beim Fußball, Handball o.Ä. vorkommen.

## Lesehinweis

Der Aufbau dieses Packages ist etwas ungewöhnlich.

Die Datei [`match.go`](match.go) enthält eine Definition des Datentyps `Match`,
wie man sie normalerweise erwarten würde.
Diese Definition besteht aus zwei Teilen:

1. Die eigentliche Struct-Definition.
2. Die Definition von Methoden, die auf dem Struct operieren.

Das Wesentliche in der Datei ist die eigentliche Typ-Definition.
Die weiteren Methoden dienen der Anschauung, wie man normalerweise vorgehen würde.
Sie werden in diesem Beispiel aber nicht weiter verwendet.

In den Beispiel- bzw. Test-Dateien werden jeweils weitere Methoden definiert.
Dies ist unüblich und bringt im Produktivcode auch nichts, weil diese Tests
für den Produktivcode gar nicht übersetzt werden.
Hier werden Methoden in Beispielen definiert, um sie direkt dort verwenden zu können.
