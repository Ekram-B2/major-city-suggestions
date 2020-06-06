package rankmanagerclient

import "testing"

func Test_createURL(t *testing.T) {
	type args struct {
		searchTerm string
		realTerm   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "NoSpaces",
			args: args{searchTerm: "tor", realTerm: "toronto"},
			want: "http://127.0.0.1:8081/determineRank?searchTerm=tor&realTerm=toronto",
		},
		{
			name: "HasSpaces",
			args: args{searchTerm: "tor", realTerm: "New York"},
			want: "http://127.0.0.1:8081/determineRank?searchTerm=tor&realTerm=New%20York",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Nothing to init to set up the test (Arrange)

			// Compute output and check if result matches the expected (Act, Assert)
			if got := createURL(tt.args.searchTerm, tt.args.realTerm); got != tt.want {
				t.Errorf("createURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
