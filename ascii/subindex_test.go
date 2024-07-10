package ascii

import (
	"reflect"
	"testing"
)

func Test_subIndexs(t *testing.T) {
	type args struct {
		inputStr string
		subStr   string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Basic test case",
			args: args{inputStr: "hello world", subStr: "o"},
			want: []int{4, 7},
		},
		{
			name: "Empty input string",
			args: args{inputStr: "", subStr: "abc"},
			want: []int{},
		},
		{
			name: "Substring not found",
			args: args{inputStr: "hello world", subStr: "xyz"},
			want: []int{},
		},
		{
			name: "Multiple occurrences",
			args: args{inputStr: "abababab", subStr: "ab"},
			want: []int{0, 2, 4, 6},
		},
		{
			name: "Overlapping occurrences",
			args: args{inputStr: "aaaaaa", subStr: "aa"},
			want: []int{0, 2, 4},
		},
		{
			name: "Case sensitivity",
			args: args{inputStr: "Hello, Hello, hello", subStr: "hello"},
			want: []int{14},
		},
		{
			name: "Edge case - Single character",
			args: args{inputStr: "aaaaa", subStr: "a"},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subIndexs(tt.args.inputStr, tt.args.subStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subIndexs() = %v, want %v", got, tt.want)
			}
		})
	}
}
