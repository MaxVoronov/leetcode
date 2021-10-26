package invert_binary_tree

import (
	"reflect"
	"testing"
)

type nodeValues []interface{}

func Test_invertTree(t *testing.T) {
	tests := []struct {
		input nodeValues
		want  nodeValues
	}{
		{input: nodeValues{4, 2, 7, 1, 3, 6, 9}, want: nodeValues{4, 7, 2, 9, 6, 3, 1}},
		{input: nodeValues{2, 1, 3}, want: nodeValues{2, 3, 1}},
		{input: nodeValues{1, 2}, want: nodeValues{1, nil, 2}},
		{input: nodeValues{}, want: nodeValues{}},
	}

	for i, tc := range tests {
		expected := createTreeFromSlice(tc.want, 0)
		got := invertTree(createTreeFromSlice(tc.input, 0))
		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, tc.input)
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
