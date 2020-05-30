package datastore

import (
	"reflect"
	"testing"
)

func Test_populateAllCities(t *testing.T) {
	type args struct {
		searchTerm string
	}
	tests := []struct {
		name    string
		args    args
		want    []LargeCity
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := populateAllCities(tt.args.searchTerm)
			if (err != nil) != tt.wantErr {
				t.Errorf("populateAllCities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("populateAllCities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_populate(t *testing.T) {
	_, err := populateAllCities("hello")
	if err != nil {
		t.Fatal(err.Error())
	}
}
