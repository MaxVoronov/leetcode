package max_depth_binary_tree

import (
	"reflect"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		input *TreeNode
		want  int
	}{
		{
			input: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: 3,
		},
		{
			input: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: &TreeNode{Val: 2},
			},
			want: 2,
		},
		{
			input: &TreeNode{
				Val:   -100,
				Left:  nil,
				Right: nil,
			},
			want: 1,
		},
		{
			input: nil,
			want:  0,
		},
	}

	for i, tc := range tests {
		got := maxDepth(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
