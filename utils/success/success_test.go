package success

import "testing"

func Test_successFunc(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "success",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := successFunc(); got != tt.want {
				t.Errorf("successFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
