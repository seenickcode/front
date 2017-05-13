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
