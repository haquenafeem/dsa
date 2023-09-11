package stack

import (
	"testing"

	"github.com/haquenafeem/dsa/utils"
)

var testCases = []struct {
	Title       string
	InputArray  []int
	OutputArray []int
}{
	{
		Title:       "Pushing [1,2,3,4] items to stack and popping them should return [4,3,2,1]",
		InputArray:  []int{1, 2, 3, 4},
		OutputArray: []int{4, 3, 2, 1},
	},
	{
		Title:       "Pushing [1,2,3,4,5] items to stack and popping them should return [5,4,3,2,1]",
		InputArray:  []int{1, 2, 3, 4, 5},
		OutputArray: []int{5, 4, 3, 2, 1},
	},
}

func Test_Stack(t *testing.T) {
	stack := New()
	for i, testCase := range testCases {
		arr := make([]int, len(testCase.InputArray))
		t.Logf("[%d] %s", i, testCase.Title)
		for i := 0; i < len(testCase.InputArray); i++ {
			stack.Push(testCase.InputArray[i])
		}

		for i := 0; i < len(testCase.InputArray); i++ {
			value, err := stack.Pop()
			if err != nil {
				t.Error(err)
			}
			arr[i] = value
		}

		if !utils.CheckIfArrayEqual(testCase.OutputArray, arr) {
			t.Fail()
		}
	}
}
