package size

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{0, "zero"},
	{5, "small"},
	{100, "huge"},
}

func TestSize(t *testing.T) {
	for _, test := range tests {
		t.Run(test.out, func(t *testing.T) {
			size := Size(test.in)
			if size != test.out {
				t.Errorf("Size(%d)=%s; want %s", test.in, size, test.out)
			}
		})
	}
}
