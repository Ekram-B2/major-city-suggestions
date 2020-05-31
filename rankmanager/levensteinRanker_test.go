package rankmanager

import "testing"

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a_greater_than_b",
			args: args{a: 6, b: -1},
			want: -1,
		},
		{
			name: "b_greater_than_a",
			args: args{a: -1, b: 6},
			want: -1,
		},
		{
			name: "a_equal_to_b",
			args: args{a: 6, b: 6},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
