package front

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContacts(t *testing.T) {
	assert := assert.New(t)
	f := NewTestFront(t)

	c1ID := os.Getenv("FRONT_TEST_CONTACT_ID")
	if len(c1ID) == 0 {
		t.Fatalf("please set env var FRONT_TEST_CONTACT_ID to run tests")
	}

	// fetch a contact
	c1, err := f.FetchContact(c1ID)
	assert.NoError(err)
	assert.NotNil(c1)

	// update contact
	p1 := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		AvatarURL   string `json:"avatar_url"`
	}{
		Name:        RandomString(),
		Description: RandomString(),
		AvatarURL:   "http://static5.businessinsider.com/image/56f7c1a4910584155c8b8dc8-1190-625/wow-donald-trump-tweets-out-fake-time-magazine-cover-naming-him-person-of-the-year.jpg",
	}
	err = f.UpdateContact(c1ID, p1)
	assert.NoError(err)

	// fetch again
	c2, err := f.FetchContact(c1ID)
	assert.NoError(err)
	assert.NotNil(c2)
	if c2 != nil {
		assert.Equal(p1.Name, c2.Name)
		assert.Equal(p1.Description, c2.Description)
		//assert.Equal(p1.AvatarURL, c2.AvatarURL) // TODO not working at the moment. I submitted a bug to the Front team
	}
}
