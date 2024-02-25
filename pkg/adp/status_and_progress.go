package adp

/* TODO
import (
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/xifanyan/adp/client"
)

type StatusAndProgressTask struct {
	client      *client.Client
	executionID string
}

func NewStatusAndProgressTask(c *client.Client, execID string) *StatusAndProgressTask {
	return &StatusAndProgressTask{
		client:      c,
		executionID: execID,
	}
}

func (t *StatusAndProgressTask) NewStatusAndProgressRequest() *http.Request {
	url := fmt.Sprintf("https://%s/%s/%s", t.client.Host, STATUSANDPROGRESS, t.executionID)

	log.Debug().Msgf("%s", url)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Auth-Username", t.client.User)
	req.Header.Set("Auth-Password", t.client.Password)

	return req
}

func (t *StatusAndProgressTask) SerializeOutput() (string, error) {
	var err error

	var req *http.Request
	var resp *http.Response

	var body []byte

	req = t.NewStatusAndProgressRequest()

	if resp, err = t.client.HTTPClient.Do(req); err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if body, err = io.ReadAll(resp.Body); err != nil {
		return "", err
	}

	return string(body), nil
}

func (t *StatusAndProgressTask) SetAsync() {}
*/
