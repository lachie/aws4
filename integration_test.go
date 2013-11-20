package aws4

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
	"github.com/nwade/aws4/assert"
)


func TestRequestWithJsonBody(t *testing.T) {
	data := strings.NewReader("{}")
	r, _ := http.NewRequest("POST", "https://dynamodb.us-east-1.amazonaws.com/", data)
	r.Header.Set("Content-Type", "application/x-amz-json-1.0")
	r.Header.Set("X-Amz-Target", "DynamoDB_20111205.ListTables")

	c, err := NewClientFromEnv()
	assert.NoError(t, err)

	resp, err := c.Do(r)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestRequestWithFormEncodedBody(t *testing.T) {
	v := make(url.Values)
	v.Set("Action", "DescribeAutoScalingGroups")

	c, err := NewClientFromEnv()
	assert.NoError(t, err)

	resp, err := c.PostForm("https://autoscaling.us-east-1.amazonaws.com/", v)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}
