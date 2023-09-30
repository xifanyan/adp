package adp

import "encoding/json"

type Request struct {
	TaskType          string `json:"taskType"`
	TaskDescription   string `json:"taskDescription"`
	TaskDisplayName   string `json:"taskDisplayName"`
	TaskConfiguration any    `json:"taskConfiguration"`
}

// JSON generates a JSON representation of the Request object.
//
// It returns a string containing the JSON representation and an error if
// there was an issue during the marshaling process.
func (req *Request) JSON() (string, error) {
	jsonBytes, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// WithDescription sets the description of the Request.
//
// description: a string representing the description of the Request.
// Returns: a pointer to the modified Request.
func (req *Request) WithDescription(description string) *Request {
	req.TaskDescription = description
	return req
}

// WithName sets the name of the Request.
//
// name: a string representing the name of the Request.
// Returns: a pointer to the modified Request.
func (req *Request) WithName(name string) *Request {
	req.TaskDisplayName = name
	return req
}
