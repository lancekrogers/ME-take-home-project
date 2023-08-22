package integration_tests

import "testing"

func TestSingleUpdate(t *testing.T) {
	// Start simulation, ID1 scheduled to be ingested 550ms later
	// 550ms ID1 v1 is "ingested", we print it as indexed
	// callbacktimeMs(400ms) later the callback fires and logs the accoundId + version
}

func TestUpdatesWithCancellation(t *testing.T) {
	// 0ms - simulation starts - ID1 scheduled to be ingested 550ms (0-1000ms random) later
	// 550ms - ID1 v1 is “ingested”, we print it as indexed
	// 650ms - ID1 v3 is “ingested”, print ID1 v3 indexed, cancel active ID1 v1 callback
	// at 950ms ensure that the  ID1 v1 callback doesn't fire
	// 1050ms - ID1 v3 callback fire
}

func TestWithMockJson(t *testing.T) {
	// Create a small test json file using a subset of data from the challenge input file

	// Test that the console log is as expected
	t.Run("Ensure the console logging is as expected for the input data", func(t *testing.T) {

	})

}
