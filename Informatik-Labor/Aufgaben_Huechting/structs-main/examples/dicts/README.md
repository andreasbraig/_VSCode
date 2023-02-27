# Wörterbücher

Hier wird ein kleines Wörterbuch mit Abfragemöglichkeiten entwickelt.
Dies geschieht in Form von zwei Packages, die aufeinander aufbauen:

* [dict1](dict1/README.md):
  Hier wird ein einfaches Datenhaltungs-Struct für Wörterbuch-Einträge
  definiert und dessen Verwendung gezeigt.
  Außerdem wird ein Wörterbuch-Datentyp definiert,
  also ein Typ für eine Sammlung solcher Einträge.

* [dict2](dict2/README.md):
    Hier wird das Datenhaltungs-Struct aus [dict1](dict1/README.md)
    um zusätzliche Funktionen und Methoden erweitert, um den Zugriff zu erleichtern.

*Anmerkung*:
Dass hier zwei einzelne Packages entwickelt werden, dient der Demonstation.
Normalerweise würde man ein Package definieren und es nach und nach verändern.
Die beiden Packages hier sollen mögliche Schritte beim Vorgehen verdeutlichen und
das Konzept schrittweise erläutern.
