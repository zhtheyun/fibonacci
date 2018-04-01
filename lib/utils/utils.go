package utils

import "math/big"

// GenerateFibonacci used to generate fibonacci string
// Params:
//    start: start from
//    start_next: second fib.
//    numbers: how many fib numbers need to generate
// Return:
//    result: fibonacci list contains {numbers}
//    start: start from {numbers}+1
//    next: second fib after {numbers}+1
func GenerateFibonacci(start big.Int, next big.Int, numbers uint64) ([]string, big.Int, big.Int) {
	var result []string
	result = make([]string, numbers)

	//The tricky thing in here is big.Int is a struct contain reference type, in this function, we need operate on the start and next which will pollute the original value. we need some way to avoid it.
	var tempStart, tempNext big.Int
	tempStart = *new(big.Int)
	tempNext = *new(big.Int)
	tempStart.SetString(start.String(), 10)
	tempNext.SetString(next.String(), 10)

	for index := range result {
		//Add start element
		result[index] = tempStart.String()
		//Compute next fib number and perform the switch
		tempStart.Add(&tempStart, &tempNext)
		tempStart, tempNext = tempNext, tempStart
	}

	return result, tempStart, tempNext
}
