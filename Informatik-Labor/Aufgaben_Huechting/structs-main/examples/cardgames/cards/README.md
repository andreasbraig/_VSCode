# Package `cards`

Hier werden die Datentypen definiert, die zu einem Kartenstapel gehören:

* ([`PlayingCard`](playingcard.go)) repräsentiert eine Spielkarte.
  Jede Karte hat jewei eine Farbe und einen Wert, z.B. *Pik Sieben* oder *Herz Ass*.
* Ein ([`Deck`](deck.go)) repräsentiert einen ganzen Kartenstapel.
  Ein Stapel ist i.W. eine Liste von Spielkarten.
* Eine ([`Hand`](hand.go)) repräsentiert die Karten,
  die ein(e) Spieler(in) auf der Hand hält.
  Auch das ist i.W. eine Liste von Spielkarten.
