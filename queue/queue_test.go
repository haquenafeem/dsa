package queue

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
		Title:       "Enqueing [1,2,3,4] items to queue and dequeing them should return [1,2,3,4]",
		InputArray:  []int{1, 2, 3, 4},
		OutputArray: []int{1, 2, 3, 4},
	},
}

func Test_Queue(t *testing.T) {
	queue := New[int]()
	for i, testCase := range testCases {
		arr := make([]int, len(testCase.InputArray))
		t.Logf("[%d] %s", i, testCase.Title)
		for i := 0; i < len(testCase.InputArray); i++ {
			queue.Enqueue(testCase.InputArray[i])
		}

		for i := 0; i < len(testCase.InputArray); i++ {
			value, err := queue.Dequeue()
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
