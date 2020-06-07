package config

import "testing"

func Test_config_GetCharDistCalculator(t *testing.T) {
	// 1. Create an instance of a system configuration object (Arrange)
	sc := SystemConfig{CharDistCalculator: "levenstein"}

	// 2. Define the expected output from the operation (Act)
	expectedOut := "levenstein"

	// 3. Perform operation to render actual output (Act)
	actualOut := sc.GetCharDistCalculator()

	// 4. Determine if the actual matches with the expected (Assert)
	if actualOut != expectedOut {
		t.Fatalf("the actual did match the expected; the actual was %s and the expected was %s", actualOut, expectedOut)
	}

}

func isSameSlice(expected, actual []string) bool {
	for index, item := range expected {
		if item != actual[index] {
			return false
		}
	}
	return true
}

func Test_config_GetMinimalKeySet(t *testing.T) {
	// 1. Create an instance of a system configuration object (Arrange)
	sc := SystemConfig{MinimalKeySet: []string{"lat, lng"}}

	// 2. Define the expected output from the operation (Act)
	expectedOut := []string{"lat, lng"}

	// 3. Perform operation to render actual output (Act)
	actualOut := sc.GetMinimalKeySet()

	// 4. Determine if the actual matches with the expected (Assert)
	if actualOut == nil {
		t.Fatalf("the actual did match the expected; the actual returned nil")
	}

	if len(actualOut) != len(expectedOut) {
		t.Fatalf("the actual did not match the expected; the actual was %v and the expected was %v", actualOut, expectedOut)
	}

	if !isSameSlice(expectedOut, actualOut) {
		t.Fatalf("the actual did match the expected; the actual was %v and the expected was %v", actualOut, expectedOut)
	}
}

func Test_config_GetDataPointType(t *testing.T) {
	// 1. Create an instance of a system configuration object (Arrange)
	sc := SystemConfig{DataPointType: "city"}

	// 2. Define the expected output from the operation (Act)
	expectedOut := "city"

	// 3. Perform operation to render actual output (Act)
	actualOut := sc.GetDataPointType()

	// 4. Determine if the actual matches with the expected (Assert)
	if actualOut != expectedOut {
		t.Fatalf("the actual did match the expected; the actual was %s and the expected was %s", actualOut, expectedOut)
	}

}

func Test_config_GetFileType(t *testing.T) {
	// 1. Create an instance of a system configuration object (Arrange)
	sc := SystemConfig{FileType: "json"}

	// 2. Define the expected output from the operation (Act)
	expectedOut := "json"

	// 3. Perform operation to render actual output (Act)
	actualOut := sc.GetFileType()

	// 4. Determine if the actual matches with the expected (Assert)
	if actualOut != expectedOut {
		t.Fatalf("the actual did match the expected; the actual was %s and the expected was %s", actualOut, expectedOut)
	}

}

func Test_config_IsRemoteClient(t *testing.T) {
	// 1. Create an instance of a system configuration object (Arrange)
	sc := SystemConfig{IsRemote: true}

	// 2. Define the expected output from the operation (Act)
	expectedOut := true

	// 3. Perform operation to render actual output (Act)
	actualOut := sc.IsRemoteClient()

	// 4. Determine if the actual matches with the expected (Assert)
	if actualOut != expectedOut {
		t.Fatalf("the actual did match the expected; the actual was %v and the expected was %v", actualOut, expectedOut)
	}

}

func testGetPath() string {
	return "files/config.test.json"
}

func Test_config_LoadConfiguration(t *testing.T) {
	// 1. Create an instance of the system configuration object (Arrange)
	sc := SystemConfig{}

	// 2. Specify what members are to be expected
	expectedDataPointType := "city"
	expectedFileType := "json"
	expectedIsRemote := false
	expectedMinimalkeyset := []string{"city", "country", "is02", "lng", "lat"}
	expectedCharDistanceCalculator := "levenstein"

	// 3. Compute the output (Act)
	newSc, err := sc.LoadConfiguration(testGetPath)

	// 4. Determine if the loaded configuration matches up with what was expected
	if err != nil {
		t.Fatalf("was not able to load configuration object")
	}

	if newSc.GetDataPointType() != expectedDataPointType {
		t.Fatalf("the actual did match the expected; the actual was %v and the expected was %v", newSc.GetDataPointType(), expectedDataPointType)
	}

	if newSc.IsRemoteClient() != expectedIsRemote {
		t.Fatalf("the actual did match the expected; the actual was %v and the expected was %v", newSc.IsRemoteClient(), expectedIsRemote)
	}

	if newSc.GetFileType() != expectedFileType {
		t.Fatalf("the actual did match the expected; the actual was %v and the expected was %v", newSc.GetFileType(), expectedFileType)
	}

	if !isSameSlice(newSc.GetMinimalKeySet(), expectedMinimalkeyset) {
		t.Fatalf("the actual did match the expected; the actual was %v and the expected was %v", newSc.GetMinimalKeySet(), expectedMinimalkeyset)
	}

	if newSc.GetCharDistCalculator() != expectedCharDistanceCalculator {
		t.Fatalf("the actual did match the expected; the actual was %v and the expected was %v", newSc.GetCharDistCalculator(), expectedCharDistanceCalculator)
	}

}
