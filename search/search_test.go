package search

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Title      string
	InputArray []int
	FindNumber int
	Output     int
}{
	{
		Title:      "Searching for number 1 in array [1,2,3,4], should return 0 as Output",
		InputArray: []int{1, 2, 3, 4},
		FindNumber: 1,
		Output:     0,
	},
	{
		Title:      "Searching for number 5 in array [1,2,3,4], should return -1 as Output",
		InputArray: []int{1, 2, 3, 4},
		FindNumber: 5,
		Output:     -1,
	},
	{
		Title:      "Searching for number 5 in array [], should return -1 as Output",
		InputArray: []int{},
		FindNumber: 5,
		Output:     -1,
	},
}

func Test_Linear(t *testing.T) {
	for i, testCase := range testCases {
		t.Log(fmt.Sprintf("[%d]", i+1), testCase.Title)
		out := Linear(testCase.InputArray, testCase.FindNumber)
		if testCase.Output != out {
			t.Log("failed")
			t.Fail()
		}
	}
}

func Test_Binary(t *testing.T) {
	for i, testCase := range testCases {
		t.Log(fmt.Sprintf("[%d]", i+1), testCase.Title)
		out := Binary(testCase.InputArray, testCase.FindNumber)
		if testCase.Output != out {
			t.Log("failed")
			t.Fail()
		}
	}
}
