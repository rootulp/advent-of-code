package main

import "fmt"

const SUBJECT_NUMBER = 7
const DIVISOR = 20201227

func PartOne(cardPublicKey int, doorPublicKey int) (int) {
	cardLoopSize := loopSize(cardPublicKey)
	doorLoopSize := loopSize(doorPublicKey)

	fmt.Printf("cardLoopSize %d\n", cardLoopSize)
	fmt.Printf("doorLoopSize %d\n", doorLoopSize)

	return encryptionKey(cardPublicKey, doorLoopSize)
}

func loopSize(publicKey int) (int) {
	val := 1
	loop := 1
	for val != publicKey {
		val *= SUBJECT_NUMBER
		val %= DIVISOR
		loop += 1
	}
	return loop - 1
}

func encryptionKey(publicKey int, loopSize int) (int) {
	var val = 1

	for loop := 0; loop < loopSize; loop += 1 {
		val *= publicKey
		val %= DIVISOR
		// fmt.Printf("loop %d val %d\n", loop, val)
	}
	return val
}
