package trail

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	body     []byte
}

func (c *Context) GetBase() *Context {
	return c
}

// Reads the body and cache it for later calls
func (c Context) Body() ([]byte, error) {
	if c.body != nil {
		return c.body, nil
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	// Restore the Body for downstream readers
	c.Request.Body = io.NopCloser(bytes.NewReader(body))

	c.body = body
	return c.body, nil
}

func (c Context) BodyNoErr() []byte {
	body, _ := c.Body()
	return body
}

// Reads the body and unmarshal it to data
func (c Context) Json(data any) error {
	body, err := c.Body()
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}

func (c Context) Write(data []byte) (int, error) {
	return c.Response.Write(data)
}

func (c Context) WriteJson(v any) (int, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return 0, err
	}

	return c.Response.Write(data)
}

func (c Context) Header(key string) string {
	return c.Request.Header.Get(key)
}
