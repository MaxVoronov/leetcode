package sum_of_left_leaves

import "testing"

type nodeValues []interface{}

func Test_sumOfLeftLeaves(t *testing.T) {
	tests := []struct {
		input nodeValues
		want  int
	}{
		{input: nodeValues{3, 9, 20, nil, nil, 15, 7}, want: 24},
		{input: nodeValues{1, 2, 3, 4, 5}, want: 4},
		{input: nodeValues{1}, want: 0},
		{input: nodeValues{}, want: 0},
	}

	for i, tc := range tests {
		sum := sumOfLeftLeaves(createTreeFromSlice(tc.input, 0))
		if tc.want != sum {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, sum)
		}
	}
}

func createTreeFromSlice(data nodeValues, pos int) *TreeNode {
	if pos > len(data)-1 {
		return nil
	}

	if val, ok := data[pos].(int); ok {
		return &TreeNode{
			Val:   val,
			Left:  createTreeFromSlice(data, 2*pos+1),
			Right: createTreeFromSlice(data, 2*pos+2),
		}
	}

	return nil
}
