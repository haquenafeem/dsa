package sort

import (
	"fmt"
	"testing"

	"github.com/haquenafeem/dsa/utils"
)

var testCases = []struct {
	Title       string
	InputArray  []int
	SortedArray []int
}{
	{
		Title:       "Given array [1,2,3,4], sorted array should be [1,2,3,4]",
		InputArray:  []int{1, 2, 3, 4},
		SortedArray: []int{1, 2, 3, 4},
	},
	{
		Title:       "Given array [4, 3, -1, -10], sorted array should be [-10, -1, 3, 4]",
		InputArray:  []int{4, 3, -1, -10},
		SortedArray: []int{-10, -1, 3, 4},
	},
	{
		Title:       "Given array [99,98,97,96], sorted array should be [96, 97, 98, 99]",
		InputArray:  []int{99, 98, 97, 96},
		SortedArray: []int{96, 97, 98, 99},
	},
	{
		Title:       "Given array [99,98,97,96], sorted array provided [96, 97, 98, 99,100] (intentional fail)",
		InputArray:  []int{99, 98, 97, 96},
		SortedArray: []int{96, 97, 98, 99, 100},
	},
}

func Test_Bubble(t *testing.T) {
	for i, testCase := range testCases {
		t.Log(fmt.Sprintf("[%d]", i+1), testCase.Title)
		out := Bubble(testCase.InputArray)
		if !utils.CheckIfArrayEqual(out, testCase.SortedArray) {
			t.Log("failed")
			t.Fail()
		}
	}
}

func Test_Selection(t *testing.T) {
	for i, testCase := range testCases {
		t.Log(fmt.Sprintf("[%d]", i+1), testCase.Title)
		out := Selection(testCase.InputArray)
		if !utils.CheckIfArrayEqual(out, testCase.SortedArray) {
			t.Log("failed")
			t.Fail()
		}
	}
}

func Test_Merge(t *testing.T) {
	for i, testCase := range testCases {
		t.Log(fmt.Sprintf("[%d]", i+1), testCase.Title)
		out := Merge(testCase.InputArray)
		if !utils.CheckIfArrayEqual(out, testCase.SortedArray) {
			t.Log("failed")
			t.Fail()
		}
	}
}
