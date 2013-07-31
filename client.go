package aws4

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var DefaultClient *Client

func init() {
	keys, err := KeysFromEnvironment()
	if err == nil {
		DefaultClient = &Client{Keys: keys}
	} else {
		log.Println("WARNING:", err)
	}
}

// Initializes and returns a Keys using the AWS_ACCESS_KEY and AWS_SECRET_KEY
// environment variables.
func KeysFromEnvironment() (keys *Keys, err error) {
	keys = &Keys{
		AccessKey: os.Getenv("AWS_ACCESS_KEY"),
		SecretKey: os.Getenv("AWS_SECRET_KEY"),
	}
	if keys.AccessKey == "" {
		err = errors.New("AWS_ACCESS_KEY not found in environment")
	}
	if keys.SecretKey == "" {
		err = errors.New("AWS_SECRET_KEY not found in environment")
	}
	return keys, err
}

// Client is like http.Client, but signs all requests using Keys.
type Client struct {
	Keys *Keys

	// The http client to make requests with. If nil, http.DefaultClient is used.
	Client *http.Client
}

// Post works like http.Post, but signs the request with Keys.
func Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error) {
	if DefaultClient == nil {
		return nil, errors.New("no DefaultClient")
	}
	return DefaultClient.Post(url, bodyType, body)
}

// PostForm works like http.PostForm, but signs the request with Keys.
func PostForm(url string, data url.Values) (resp *http.Response, err error) {
	if DefaultClient == nil {
		return nil, errors.New("no DefaultClient")
	}
	return DefaultClient.PostForm(url, data)
}

func (c *Client) client() *http.Client {
	if c.Client == nil {
		return http.DefaultClient
	}
	return c.Client
}

func (c *Client) Do(req *http.Request) (resp *http.Response, err error) {
	Sign(c.Keys, req)
	return c.client().Do(req)
}

func (c *Client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) Head(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	return c.Do(req)
}

func (c *Client) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}
