package aufgabe6

import "fmt"

func ExampleConnection_IsBefore() {
	first := Connection{
		Origin:          "Mannheim",
		Destination:     "Frankfurt",
		DepartureHour:   13,
		DepartureMinute: 15,
		ArrivalHour:     14,
		ArrivalMinute:   24,
	}
	second_bad := Connection{
		Origin:          "Frankfurt",
		Destination:     "Kassel",
		DepartureHour:   14,
		DepartureMinute: 21,
		ArrivalHour:     16,
		ArrivalMinute:   55,
	}
	second_good := Connection{
		Origin:          "Frankfurt",
		Destination:     "Kassel",
		DepartureHour:   15,
		DepartureMinute: 19,
		ArrivalHour:     17,
		ArrivalMinute:   34,
	}

	fmt.Println(first.IsBefore(second_bad))
	fmt.Println(first.IsBefore(second_good))

	// Output:
	// false
	// true
}
