package utils

import "math/big"
import "testing"
import "reflect"

func TestGenerateFibonacci(t *testing.T) {
	fibResultSet := []string{
		"0",
		"1",
		"1",
		"2",
		"3",
		"5",
		"8",
		"13",
		"21",
		"34",
		"55",
		"89",
		"144",
		"233",
		"377",
		"610",
		"987",
		"1597",
		"2584",
		"4181",
		"6765",
		"10946",
	}

	var testCases = []struct {
		start               big.Int
		next                big.Int
		numbers             uint64
		expectedResultData  []string
		expectedResultStart string
		expectedResultNext  string
	}{
		{
			*big.NewInt(0),
			*big.NewInt(1),
			1,
			fibResultSet[0:1],
			fibResultSet[1],
			fibResultSet[2],
		},
		{
			*big.NewInt(0),
			*big.NewInt(1),
			2,
			fibResultSet[0:2],
			fibResultSet[2],
			fibResultSet[3],
		},
		{
			*big.NewInt(0),
			*big.NewInt(1),
			5,
			fibResultSet[0:5],
			fibResultSet[5],
			fibResultSet[6],
		},
		{
			*big.NewInt(0),
			*big.NewInt(1),
			10,
			fibResultSet[0:10],
			fibResultSet[10],
			fibResultSet[11],
		},
		{
			*big.NewInt(0),
			*big.NewInt(1),
			20,
			fibResultSet[0:20],
			fibResultSet[20],
			fibResultSet[21],
		},

		{
			*big.NewInt(1),
			*big.NewInt(2),
			3,
			fibResultSet[2:5],
			fibResultSet[5],
			fibResultSet[6],
		},
	}

	for _, testcase := range testCases {
		actualResultData, start, next := GenerateFibonacci(testcase.start, testcase.next, testcase.numbers)
		if !StringSliceReflectEqual(actualResultData, testcase.expectedResultData) {
			t.Errorf("Result data did not match expectation. expect: %s, actual: %s", testcase.expectedResultData, actualResultData)
		}

		if start.String() != testcase.expectedResultStart {
			t.Errorf("Result start did not match expectation. expect: %s, actual: %s", testcase.expectedResultStart, start.String())

		}

		if next.String() != testcase.expectedResultNext {
			t.Errorf("Result start did not match expectation. expect: %s, actual: %s", testcase.expectedResultNext, next.String())

		}

	}
}

func ByteSliceReflectEqual(a, b []byte) bool {
	return reflect.DeepEqual(a, b)
}
func StringSliceReflectEqual(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

func BigintSliceReflectEqual(a, b []big.Int) bool {
	return reflect.DeepEqual(a, b)
}
