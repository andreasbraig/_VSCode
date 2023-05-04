# Aufgabe: Aufbau einer Shopping-Datenbank

Definieren Sie Datentypen für die Einträge eines Online-Shops.
Denken Sie dabei über verschiedene Kategorien nach:
Es kann z.B. sowohl Produkte als auch Kunden geben.

Ein Datensatz zu einem einzelnen Produkt könnte z.B. enthalten:

* Name des Produkts
* Verkaufspreis
* Einkaufspreis
* Lagerbestand
* Kategorien/Schlagwörter (z.B. "Kosmetik", "Technik", "Entertainment" o.Ä.)

Ein Kunde kann z.B. folgendes enthalten:

* Name
* Kundennummer
* Umsatz der letzten 12 Monate
* Liste der zuletzt gekauften Produkte

Die gesamte Datenbank enthält dann sowohl eine Liste von Kunden als auch eine Liste
von Produkten.

Definieren Sie Funktionen, mit denen auf derartige Daten zugegriffen werden kann.
Dazu gehören einfache, aber auch komplexere Abfragen, wie z.B.:

* Name zu einer gegebenen Kundennummer
* Alle Produkte einer bestimmten Kategorie
* Alle Produkte mit niedrigem Lagerbestand
* Die häufigste gekaufte Produktkategorie eines Kunden
* Die zweithäufigste gekaufte Produktkategorie eines Kunden
* Empfehlungen für Kunden aufgrund der Käufe anderer Kunden

*Hinweis*: Fangen Sie langsam an und definieren Sie zunächst einfache Datentypen für
Produkt und/oder Kunde und implementieren Sie einfache Abfragen darauf, bevor Sie
weitermachen.
