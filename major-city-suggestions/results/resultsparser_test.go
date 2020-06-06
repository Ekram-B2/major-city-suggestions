package results

import "testing"

func Test_isAMember(t *testing.T) {
	type args struct {
		key        string
		properties []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Success",
			args: args{key: "country", properties: []string{"country", "iso2"}},
			want: true,
		},
		{
			name: "Failure",
			args: args{key: "country", properties: []string{"lat", "iso2"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Nothing to init for Arrange set (Arrange)

			// Apply operation to get output and determine if actual matched expected (Act, Assert)
			if got := isAMember(tt.args.key, tt.args.properties); got != tt.want {
				t.Errorf("isAMember() = %v, want %v", got, tt.want)
			}
		})
	}
}
