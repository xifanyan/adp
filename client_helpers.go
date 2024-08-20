package adp

import "encoding/json"

func (c *Client) ListEntities(opts ...func(*ListEntitiesConfiguration)) ([]Entity, error) {
	var err error
	var resp *Response

	req := NewRequest().ListEntities(opts...)

	if resp, err = c.Send(req); err != nil {
		return nil, err
	}

	var meta ListEntitiesExecutionMetaData

	if err = json.Unmarshal(resp.ExecutionMetaData, &meta); err != nil {
		return nil, err
	}

	js := string(meta.AdpEntitiesJSONOutput)
	UnquoteJSONOutput(&js)

	var res = []Entity{}

	// when no entity was found
	if err = json.Unmarshal([]byte(js), &res); err != nil {
		return res, nil
	}

	return res, err
}
