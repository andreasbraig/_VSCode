// Thema der Aufgabe: Bäume
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe5

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
	// TODO

	if oldString[0] == newString[0] {
		for _, bucket := range list {
			for i, e := range bucket {
				if e == oldString {
					bucket[i] = newString
				}
			}
		}
	} else {
		list.Add(newString)
		for bi, bucket := range list {
			for i, e := range bucket {
				if e == oldString {

					RemoveFromBucket(&list[bi], i)

				}
			}
		}

	}

}

func FindPosInBucket(bucket []string, oldstring string) int {
	for i, e := range bucket {
		if e == oldstring {
			return i
		}
	}
	return -1
}

func RemoveFromBucket(list *[]string, index int) {

	if len(*list) <= 1 || index < 0 {
		return
	}
	if len(*list) > 2 {
		copyofList := *list
		result := copyofList[:index]
		result = append(result, copyofList[index+1:]...)

		*list = result
	}

	if len(*list) == 2 {
		if index == 0 {
			copyofList := *list
			result := []string{}
			result = append(result, copyofList[1])

			*list = result
		}
		if index == 1 {
			copyofList := *list
			result := []string{}
			result = append(result, copyofList[0])

			*list = result
		}
	}

}

func (list *BucketList) RemoveBucket(bindex int)
