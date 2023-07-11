// Thema der Aufgabe: Bäume
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe5

import "strings"

/* Aufgabenstellung:
 * In der Datei bucketlist.go ist ein Datentyp für einen neuen Listen-Datentyp definiert.
 * Eine BucketList enthält 26 sogenannte Buckets, die jeweils eine Liste von Strings
 * speichern. Für jeden Buchstaben im Alphabet ein Bucket.
 * Wird ein String hinzugefügt, so wird er ans Ende des entsprechenden Buckets
 * angehängt, eine weitere Sortierung innerhalb der Buckets geschieht nicht.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode Replace().
 * Die Methode soll zwei Strings oldString und newString erwarten und den alten durch
 * den neuen ersetzen.
 *
 * Kommt oldString nicht vor, soll nichts passieren.
 * Falls oldString und newString in den selben Bucket gehören, soll newString an der
 * selben Stelle eingesetzt werden, an der oldString gestanden hat.
 */

// Ersetzt oldstring durch newString in der Liste.
func (list BucketList) Replace(oldString, newString string) {
	oldBucket := strings.ToLower(oldString)[0] - 'a'

	oldPosInBucket := -1
	for i, v := range list[oldBucket] {
		if v == oldString {
			oldPosInBucket = i
		}
	}
	if oldPosInBucket == -1 {
		return
	}
	newBucket := strings.ToLower(newString)[0] - 'a'
	if oldBucket == newBucket {
		list[oldBucket][oldPosInBucket] = newString
	} else {
		beginning := list[oldBucket][:oldPosInBucket]
		end := list[oldBucket][oldPosInBucket+1:]
		list[oldBucket] = append(beginning, end...)
		list.Add(newString)
	}
}
