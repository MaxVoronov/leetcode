package min_stack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinStack_Case1(t *testing.T) {
	minStack := Constructor()
	funcs := []string{"push", "push", "push", "getMin", "pop", "top", "getMin"}
	args := [][]int{{-2}, {0}, {-3}, {}, {}, {}, {}}
	outs := []interface{}{nil, nil, nil, -3, nil, 0, -2}

	if err := checkCases(&minStack, funcs, args, outs); err != nil {
		t.Fatal(err)
	}
}

func TestMinStack_Case2(t *testing.T) {
	minStack := Constructor()
	funcs := []string{"push", "push", "push", "top", "pop", "getMin", "pop", "getMin", "pop", "push", "top", "getMin", "push", "top", "getMin", "pop", "getMin"}
	args := [][]int{{2147483646}, {2147483646}, {2147483647}, {}, {}, {}, {}, {}, {}, {2147483647}, {}, {}, {-2147483648}, {}, {}, {}, {}}
	outs := []interface{}{nil, nil, nil, 2147483647, nil, 2147483646, nil, 2147483646, nil, nil, 2147483647, 2147483647, nil, -2147483648, -2147483648, nil, 2147483647}

	if err := checkCases(&minStack, funcs, args, outs); err != nil {
		t.Fatal(err)
	}
}

func checkCases(minStack *MinStack, funcs []string, args [][]int, outs []interface{}) error {
	fnMapping := map[string]interface{}{
		"getMin": minStack.GetMin,
		"push":   minStack.Push,
		"pop":    minStack.Pop,
		"top":    minStack.Top,
	}

	for i, name := range funcs {
		fn := reflect.ValueOf(fnMapping[name])
		if fnNum, argsNum := fn.Type().NumIn(), len(args[i]); fnNum != argsNum {
			return fmt.Errorf(`function "%s" has not same nums of arguments: expected: %d, got: %d`, name, fnNum, argsNum)
		}

		inArgs := make([]reflect.Value, 0, len(args[i]))
		for _, val := range args[i] {
			inArgs = append(inArgs, reflect.ValueOf(val))
		}

		fnRes := fn.Call(inArgs)

		var got interface{}
		if len(fnRes) > 0 {
			got = fnRes[0].Interface()
		}

		if expected := outs[i]; !reflect.DeepEqual(expected, got) {
			return fmt.Errorf(`call "%s" on step %d return invalid result: expected %+v, got: %+v`, name, i+1, expected, got)
		}
	}

	return nil
}
