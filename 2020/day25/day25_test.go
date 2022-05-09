package main

import "testing"

func TestPartOne(t *testing.T) {
	type testCase struct {
		cardPublicKey int
		doorPublicKey int
		encryptionKey int
	}

	tests := []testCase {
		{5764801, 17807724, 14897079}, // example
		// {10943862, 12721030, 0}, // input
	}

	for _, test := range tests {
		got := PartOne(test.cardPublicKey, test.doorPublicKey)

		if got != test.encryptionKey {
			t.Errorf("PartOne(%d, %d) got %d want %d", test.cardPublicKey, test.doorPublicKey, got, test.encryptionKey)
		}
	}
}
