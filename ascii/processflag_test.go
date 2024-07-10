package ascii

import "testing"

func TestProcessColorFlag(t *testing.T) {
	type parameters struct {
		colorFlag string
		arg       []string
	}
	type want struct {
		inputString, subString string
	}
	tests := []struct {
		name     string
		args     parameters
		expected want
	}{
		{
			name:     "test with all arguments",
			args:     parameters{colorFlag: "blue", arg: []string{"--color=red", "ell", "Hello"}},
			expected: want{subString: "ell", inputString: "Hello"},
		},
		{
			name:     "test without substring",
			args:     parameters{colorFlag: "#ff0000", arg: []string{"#ff0000", "Hello"}},
			expected: want{inputString: "Hello", subString: "Hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subString, inputString = "", "" // clear variables before each test

			ProcessColorFlag(tt.args.colorFlag, tt.args.arg)
			if inputString != tt.expected.inputString {
				t.Errorf("ProcessColorFlag() inputString = %v, expected %v,", inputString, tt.expected.inputString)
			}
			if subString != tt.expected.subString {
				t.Errorf("ProcessColorFlag() subString =  %v, expected %v,", subString, tt.expected.subString)
			}
		})
	}
}
