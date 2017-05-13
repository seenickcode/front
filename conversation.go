package front

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Conversation struct {
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
	ID       string `json:"id"`
	Subject  string `json:"subject"`
	Status   string `json:"status"`
	Assignee struct {
		Links struct {
			Self    string `json:"self"`
			Related struct {
				Inboxes       string `json:"inboxes"`
				Conversations string `json:"conversations"`
			} `json:"related"`
		} `json:"_links"`
		ID          string `json:"id"`
		Email       string `json:"email"`
		Username    string `json:"username"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		IsAdmin     bool   `json:"is_admin"`
		IsAvailable bool   `json:"is_available"`
	} `json:"assignee"`
	Recipient struct {
		Links struct {
			Related struct {
				Contact string `json:"contact"`
			} `json:"related"`
		} `json:"_links"`
		Handle string `json:"handle"`
		Role   string `json:"role"`
	} `json:"recipient"`
	Tags []struct {
		Links struct {
			Self    string `json:"self"`
			Related struct {
				Conversations string `json:"conversations"`
			} `json:"related"`
		} `json:"_links"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
	LastMessage struct {
		Links struct {
			Self    string `json:"self"`
			Related struct {
				Conversation     string `json:"conversation"`
				MessageRepliedTo string `json:"message_replied_to"`
			} `json:"related"`
		} `json:"_links"`
		ID        string  `json:"id"`
		Type      string  `json:"type"`
		IsInbound bool    `json:"is_inbound"`
		CreatedAt float64 `json:"created_at"`
		Blurb     string  `json:"blurb"`
		Author    struct {
			Links struct {
				Self    string `json:"self"`
				Related struct {
					Inboxes       string `json:"inboxes"`
					Conversations string `json:"conversations"`
				} `json:"related"`
			} `json:"_links"`
			ID          string `json:"id"`
			Email       string `json:"email"`
			Username    string `json:"username"`
			FirstName   string `json:"first_name"`
			LastName    string `json:"last_name"`
			IsAdmin     bool   `json:"is_admin"`
			IsAvailable bool   `json:"is_available"`
		} `json:"author"`
		Recipients []struct {
			Links struct {
				Related struct {
					Contact string `json:"contact"`
				} `json:"related"`
			} `json:"_links"`
			Handle string `json:"handle"`
			Role   string `json:"role"`
		} `json:"recipients"`
		Body        string `json:"body"`
		Text        string `json:"text"`
		Attachments []struct {
			Filename    string `json:"filename"`
			URL         string `json:"url"`
			ContentType string `json:"content_type"`
			Size        int    `json:"size"`
			Metadata    struct {
				IsInline bool   `json:"is_inline"`
				Cid      string `json:"cid"`
			} `json:"metadata"`
		} `json:"attachments"`
		Metadata struct {
		} `json:"metadata"`
	} `json:"last_message"`
	CreatedAt float64 `json:"created_at"`
}

// FetchConversation .
func (f *Front) FetchConversation(id string) (c *Conversation, err error) {
	endpoint := fmt.Sprintf("%s%s/conversations/%s", FrontHostname, FrontBaseEndpoint, id)
	data, status, err := httpCallWithAuthToken("GET", endpoint, bytes.NewBuffer([]byte{}), f.jwtToken)
	if err != nil {
		return
	}
	if status != http.StatusOK {
		err = fmt.Errorf("posting Front conversation message failed with status %v: %v", status, string(data))
		return
	}
	c = &Conversation{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return
	}
	return
}
