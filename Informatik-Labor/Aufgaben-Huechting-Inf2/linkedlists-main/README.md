# Beispielcode/Aufgaben zu verketteten Listen

## Aufgaben zur vorhandenen Liste von Zahlen

Im Package `dllist_int` ist eine Implementierung einer doppelt verketteten Liste
vorgegeben, deren Element Zahlen als Werte enthalten.
Es sind noch einige Implementierungen von Methoden/Funktionen zu vervollständigen:

* Die Funktion `swapNodes` erwartet zwei Knoten aus einer Liste und soll sie vertauschen.
* Die Methoden `PushBack`, `PushFront`, `PopBack` und `PopFront` aus dem Datentyp `DLList`
  sollen die entsprechenden Operationen auf der Liste ausführen.
* Die Methode `Get`aus dem Datentyp `DLList` soll den Wert des Elements an einer
  bestimmten Position zurückgeben.
* Die Methode `Swap`aus dem Datentyp `DLList` soll zwei Knoten aus der Liste vertauschen.

*Hinweis*: Im Package `dllist_int_solution' ist eine Lösung zu finden.

## Erweiterung des Datentyps

Fügen Sie weitere Methoden oder Funktionen hinzu,
die mehr Abfragen oder Operationen ermöglichen.
Fügen Sie auch entsprechende Tests hinzu, um die Funktionalität zu überprüfen.

**Beispiele für `DLList`:**

* `Remove` erwartet einen Index und entfernt dieses Element aus der Liste, falls es existiert. √
* `Insert` erwartet einen Index und einen Wert und fügt ein neues Element an dieser Position ein.√
* `InsertSorted` erwartet einen Wert und fügt ein neues Element an der richtigen Position ein.
* `Reverse` dreht die Reihenfolge der Elemente in der Liste um.
* `Replace` erwartet einen Index und einen Wert und ersetzt den Wert des Elements an dieser Position. √

**Beispiele für `DLNode`:**

* `Find` erwartet einen Wert und liefert den entsprechenden Element-Pointer, falls solch ein Element existiert. √
* `FindLast` erwartet einen Wert und liefert den Pointer auf das letzte Element mit diesem Wert.
* `FindAll` erwartet einen Wert und liefert eine Liste mit allen Elementen mit diesem Wert.
* `FindIndex` erwartet einen Wert und liefert den Index des ersten Elements mit diesem Wert. √

## Erstellung einer Liste für einen komplexeren Datentyp

Erstellen Sie auf Basis des Packages `dllist_int` ein Package für einen Listen-Datentyp,
dessen Elemente nicht nur Zahlen, sondern komplexere Werte enthalten.

**Einfache Beispiele/Ideen für Daten anstelle von Zahlen:**

* Einträge in einem Telefonbuch mit Namen, Telefon, E-Mail, ...
* Artikel mit Artikelnummer, Bezeichnung, Preis, ...
* Kinofilme mit Titel, Erscheinungsjahr, Genre, Bewertungen ...
* Personen in einem sozialen Netzwerk mit Name sowie jeweils einer Liste von Freunden.

*Hinweis:* Achten Sie darauf, Ihren Datentyp zunächst möglichst einfach zu halten.
Auch wenn Sie vielleicht sofort eine Reihe an Ideen haben, sollten Sie zunächst
nur das Minumum umsetzen, um die Funktionalität der Liste zu testen.
Erweitern Sie Ihren Datentyp dann schrittweise um weitere Eigenschaften.

*Hinweis:* Sie können die Lösung aus dem Package `dllist_int_solution` als Vorlage
verwenden, um die Implementierung zu vereinfachen.
