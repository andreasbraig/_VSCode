# Minimalistisches Wörterbuch

Hier wird ein Datentyp für Einträge eines Wörterbuchs definiert.
Ein solcher Eintrag besteht dabei aus einem deutschen und einem englischen Wort,
die jeweils ein String sind.
Zusätzlich wird ein Datentyp für das gesamte Wörterbuch definiert.
Ein solches Wörterbuch ist i.W. eine Liste von Einträgen.

Zu beiden Datentypen gibt es Beispiele in Test-Form.

## Datei-Überblick

* [entry.go](entry.go):
  Hier wird der Datentyp für Einträge als Struct definiert.
* [example_entry_test.go](example_entry_test.go):
  Hier gibt es Beispiele, wie man `Entry`-Objekte erzeugt und verwendet.
* [dict.go](dict.go):
  Hier wird der Datentyp für das Wörterbuch als Struct definiert.
* [example_dict_test.go](example_dict_test.go):
  Hier gibt es Beispiele, wie man `Dict`-Objekte erzeugt und verwendet.
  
## Lesehinweis

Schauen Sie sich zuerst den Datentyp `Entry` in [entry.go](entry.go) an.
Diese Definition und die Beispiele in [example_entry_test.go](example_entry_test.go)
sind die Grundlage für den Wörterbuch-Typ.
