package rankmanager

import "testing"

func Test_rankmanager_min(t *testing.T) {
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
			name: "AGreaterThanB",
			args: args{a: 6, b: -1},
			want: -1,
		},
		{
			name: "BGreaterThanA",
			args: args{a: -1, b: 6},
			want: -1,
		},
		{
			name: "AEqualToB",
			args: args{a: 6, b: 6},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Nothing to init to set up the test (Arrange)

			// Peform operation and check to see if output matches the expected (Act, Assert)
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
