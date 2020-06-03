package datastore

import (
	"testing"
)

func Test_getAllCities(t *testing.T) {
	fileManager := JSONFileManager{}
	_, err := fileManager.getAllCities()
	if err != nil {
		t.Fatal("Break")
	}

}
