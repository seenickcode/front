package front

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConversations(t *testing.T) {
	assert := assert.New(t)
	f := NewTestFront(t)

	c1ID := os.Getenv("FRONT_TEST_CONVERSATION_ID")
	if len(c1ID) == 0 {
		t.Fatalf("please set env var FRONT_TEST_CONVERSATION_ID to run tests")
	}

	// fetch a conversation
	c1, err := f.FetchConversation(c1ID)
	assert.NoError(err)
	assert.NotNil(c1)

	// update conversation
	p1 := struct {
		Status string `json:"status"`
	}{
		Status: ConversationStatusArchived,
	}
	err = f.UpdateConversation(c1ID, p1)
	assert.NoError(err)

	// fetch again
	c2, err := f.FetchConversation(c1ID)
	assert.NoError(err)
	assert.NotNil(c2)
	if c2 != nil {
		assert.Equal(p1.Status, c2.Status)
	}

	// update conversation again
	p2 := struct {
		AssigneeID *string `json:"assignee_id"`
		Status     string  `json:"status"`
	}{
		AssigneeID: nil,
		Status:     ConversationStatusOpen,
	}
	err = f.UpdateConversation(c2.ID, p2)
	assert.NoError(err)

	// fetch again
	c3, err := f.FetchConversation(c2.ID)
	assert.NoError(err)
	assert.NotNil(c3)
	// NOTE: for some reason, when updating, if you nil out the assignee ID, even if you
	// also set it to 'open', the updated conversation will be automatically switched
	// to status 'unassigned'. So it seems 'unassigned' = 'open' + no assignee ID
	if c3 != nil {
		assert.Equal("", c3.Assignee.ID)
		assert.Equal(ConversationStatusUnassigned, c3.Status)
	}
}

func NewTestFront(t *testing.T) *Front {
	jwtToken := os.Getenv("FRONT_JWT_TOKEN")
	if len(jwtToken) == 0 {
		t.Fatalf("please set env var FRONT_JWT_TOKEN to run tests")
	}
	f, err := New(jwtToken)
	if err != nil {
		t.Fatal(err)
	}
	return f
}
