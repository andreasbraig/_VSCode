package aufgabe5

import "fmt"

func ExampleBucketList_Replace_firstOfBucketSameBucket_original() {
	b1 := NewBucketList()

	b1.Add("Hallo")
	b1.Add("Haus")
	b1.Add("Auto")
	b1.Add("Fahrrad")

	b1.Replace("Hallo", "Hi")
	fmt.Println(b1)

	// Output:
	// a: [Auto]
	// f: [Fahrrad]
	// h: [Hi Haus]
}

func ExampleBucketList_Replace_lastOfBucketSameBucket_original() {
	b1 := NewBucketList()

	b1.Add("Hallo")
	b1.Add("Haus")
	b1.Add("Auto")
	b1.Add("Fahrrad")

	b1.Replace("Haus", "Halle")
	b1.Replace("Auto", "Abfall")
	fmt.Println(b1)

	// Output:
	// a: [Abfall]
	// f: [Fahrrad]
	// h: [Hallo Halle]
}

func ExampleBucketList_Replace_firstOfBucketDifferentBucket_original() {
	b1 := NewBucketList()

	b1.Add("Hallo")
	b1.Add("Haus")
	b1.Add("Auto")
	b1.Add("Fahrrad")

	b1.Replace("Hallo", "Grüezi")
	fmt.Println(b1)

	// Output:
	// a: [Auto]
	// f: [Fahrrad]
	// g: [Grüezi]
	// h: [Haus]
}

func ExampleBucketList_Replace_lastOfBucketDifferentBucket_original() {
	b1 := NewBucketList()

	b1.Add("Hallo")
	b1.Add("Haus")
	b1.Add("Auto")
	b1.Add("Fahrrad")

	b1.Replace("Haus", "Gebäude")
	b1.Replace("Auto", "Skateboard")
	fmt.Println(b1)

	// Output:
	// f: [Fahrrad]
	// g: [Gebäude]
	// h: [Hallo]
	// s: [Skateboard]
}
