package fail

import "testing"

func Test_failFunc(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "fail",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := failFunc(); got != tt.want {
				t.Errorf("failFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
