package adp

import "testing"

// TestRequest is a function that tests the Request struct.
//
// It creates a new instance of Request, sets its description and name using
// the WithDescription and WithName methods, and then checks that the
// TaskDescription and TaskDisplayName fields are set correctly. If the
// values are not as expected, it will log an error.
func TestRequest(t *testing.T) {
	req := &Request{}
	req.WithDescription("desc").WithName("name")
	if req.TaskDescription != "desc" || req.TaskDisplayName != "name" {
		t.Errorf("Unexpected request values: %v", req)
	}
}
