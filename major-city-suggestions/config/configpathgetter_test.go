package config

import (
	"os"
	"testing"
)

func Test_config_getDefaultConfigPathLocal(t *testing.T) {
	// 1. Store the value of `LOCAL` env variable (Arrange)
	temp := os.Getenv("DEVELOPMENT")

	// 2. Set the value of `LOCAL` env (Arrange)
	err := os.Setenv("DEVELOPMENT", "1")
	if err != nil {
		t.Fatalf("was not able to perform the necessary step of setting the `LOCAL` env variable")
	}
	// 3. Define the expected output from the operation (Act)
	expectedOut := "/files/config.json"

	// 4. Perform operation to render actual output (Act)
	actualOut := getDefaultConfigPath()

	// 5. Determine if the actual matches with the expected (Assert)
	if actualOut != expectedOut {
		t.Fatalf("the actual did match the expected; the actual was %s and the expected was %s", actualOut, expectedOut)
	}

	// 6. Tear down logic - return LOCAL env back to expected form
	os.Setenv("LOCAL", temp)
}
