package config

import "testing"

func testGetPath() string {
	return "files/config.test.json"
}

func Test_config_LoadConfiguration(t *testing.T) {

	// 1. Set up expectations to compare test results againsts
	wantDataPoint := "city"
	wantRemoteStatus := false
	// 2. Compute the output (Act)
	newConfig, err := LoadConfiguration(testGetPath)

	// 3. Determine if the loaded configuration matches up with what was expected (Assert)
	if err != nil {
		t.Fatalf("was not able to load configuration object")
	}

	if newConfig.DataPointType != wantDataPoint {
		t.Fatalf("the actual did not match the expected; the actual was %v and the expected was %v", newConfig.DataPointType, wantDataPoint)
	}

	if newConfig.IsRemote != wantRemoteStatus {
		t.Fatalf("the actual did not match the expected; the actual was %v and the expected was %v", newConfig.IsRemote, wantRemoteStatus)
	}

}
