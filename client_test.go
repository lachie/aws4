package aws4

import (
	"testing"
	"os"
	"github.com/nwade/aws4/assert"
)

func TestCreatingClientFromEnv(t *testing.T) {
	env := map[string]string{
		"AWS_ACCESS_KEY": "some-access-key",
		"AWS_SECRET_KEY": "some-secret-key",
	}

	for envVar, newVal := range env {
		oldVal := os.Getenv(envVar)
		os.Setenv(envVar, newVal)
		if oldVal != "" {
			defer os.Setenv(envVar, oldVal)
		}
	}

	c, err := NewClientFromEnv()
	assert.NoError(t, err)
	assert.NotNil(t, c.Keys)
	assert.Equal(t, c.Keys.AccessKey, "some-access-key")
	assert.Equal(t, c.Keys.SecretKey, "some-secret-key")
}

func TestCreatingClientWithKeys(t *testing.T) {
	c := NewClient("some-access-key", "some-secret-key")

	assert.NotNil(t, c.Keys)
	assert.Equal(t, c.Keys.AccessKey, "some-access-key")
	assert.Equal(t, c.Keys.SecretKey, "some-secret-key")
}
