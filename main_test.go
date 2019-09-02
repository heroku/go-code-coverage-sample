package main

import "testing"

func Test_run(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "simple",
			want: "2 * 3 = 6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
