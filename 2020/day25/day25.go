package main

import "fmt"

const SUBJECT_NUMBER = 7
const DIVISOR = 20201227

func main() {
	fmt.Printf("Startind day25...\n")

	partOne := PartOne(10943862, 12721030)
	fmt.Printf("PartOne: %d\n", partOne)
}

func PartOne(cardPublicKey int, doorPublicKey int) (int) {
	doorLoopSize := loopSize(doorPublicKey)
	return encryptionKey(cardPublicKey, doorLoopSize)
}

func loopSize(publicKey int) (int) {
	val := 1
	loop := 0
	for val != publicKey {
		loop += 1
		val *= SUBJECT_NUMBER
		val %= DIVISOR
	}
	return loop
}

func encryptionKey(publicKey int, loopSize int) (int) {
	var val = 1

	for loop := 0; loop < loopSize; loop += 1 {
		val *= publicKey
		val %= DIVISOR
	}
	return val
}
