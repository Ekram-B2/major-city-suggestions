package config

import (
	"os"
	"testing"
)

func Test_config_getDefaultConfigPathLocal(t *testing.T) {
	// 1. Store the value of `SYSTEM_BUILD` env variable (Arrange)
	temp := os.Getenv("SYSTEM_BUILD")

	// 2. Set the value of `SYSTEM_BUILD` env (Arrange)
	err := os.Setenv("SYSTEM_BUILD", "1")
	if err != nil {
		t.Fatalf("was not able to perform the necessary step of setting the `SYSTEM_BUILD` env variable")
	}
	// 3. Define the expected output from the operation (Act)
	expectedOut := "config/files/config.development.json"

	// 4. Perform operation to render actual output (Act)
	actualOut := getConfigPathDefault()

	// 5. Determine if the actual matches with the expected (Assert)
	if actualOut != expectedOut {
		t.Fatalf("the actual did not match the expected; the actual was %s and the expected was %s", actualOut, expectedOut)
	}

	// 6. Tear down logic - return LOCAL env back to expected form
	os.Setenv("LOCAL", temp)
}
