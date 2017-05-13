package front

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type EventPreview struct {
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	EmittedAt float64 `json:"emitted_at"`
	Source    struct {
		Meta struct {
			Type string `json:"type"`
		} `json:"_meta"`
		Links struct {
			Self string `json:"self"`
		} `json:"_links"`
		ID string `json:"id"`
	} `json:"source"`
	Target struct {
		Meta struct {
			Type string `json:"type"`
		} `json:"_meta"`
		Links struct {
			Self string `json:"self"`
		} `json:"_links"`
		ID string `json:"id"`
	} `json:"target"`
	Conversation struct {
		Links struct {
			Self    string `json:"self"`
			Related struct {
				Events    string `json:"events"`
				Followers string `json:"followers"`
				Messages  string `json:"messages"`
				Comments  string `json:"comments"`
				Inboxes   string `json:"inboxes"`
			} `json:"related"`
		} `json:"_links"`
		ID string `json:"id"`
	} `json:"conversation"`
}

func NewEventPreviewFromRequest(r *http.Request) (e *EventPreview, err error) {
	body, err := ioutil.ReadAll(r.Body)
	var obj EventPreview
	if len(string(body)) > 0 {
		if err := json.Unmarshal(body, &obj); err != nil {
			return nil, err
		}
		e = &obj
	}
	return
}
