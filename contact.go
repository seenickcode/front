package front

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Contact struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	AvatarURL   string   `json:"avatar_url"`
	IsSpammer   bool     `json:"is_spammer"`
	Links       []string `json:"links"`
	Handles     []struct {
		Handle string `json:"handle"`
		Source string `json:"source"`
	} `json:"handles"`
	Groups []struct {
		Links struct {
			Self    string `json:"self"`
			Related struct {
				Contacts string `json:"contacts"`
			} `json:"related"`
		} `json:"_links"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"groups"`
}

// FetchContact .
func (f *Front) FetchContact(id string) (c *Contact, err error) {
	endpoint := fmt.Sprintf("%s%s/contacts/%s", FrontHostname, FrontBaseEndpoint, id)
	data, status, err := httpCallWithAuthToken("GET", endpoint, bytes.NewBuffer([]byte{}), f.jwtToken)
	if err != nil {
		return
	}
	if status != http.StatusOK {
		err = fmt.Errorf("fetching Front conversation failed with status %v: %v", status, string(data))
		return
	}
	c = &Contact{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return
	}
	return
}

// UpdateContact .
func (f *Front) UpdateContact(id string, params interface{}) (err error) {
	endpoint := fmt.Sprintf("%s%s/contacts/%s", FrontHostname, FrontBaseEndpoint, id)
	paramBytes, err := json.Marshal(params)
	if err != nil {
		return
	}
	data, status, err := httpCallWithAuthToken("PATCH", endpoint, bytes.NewBuffer(paramBytes), f.jwtToken)
	if err != nil {
		return
	}
	if status != http.StatusNoContent {
		err = fmt.Errorf("updating Front conversation failed with status %v: %v", status, string(data))
		return
	}
	return
}
