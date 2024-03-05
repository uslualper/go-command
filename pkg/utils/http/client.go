package http

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

type Client struct {
	HTTPClient *fasthttp.Client
	Request    *fasthttp.Request
	Response   *fasthttp.Response
	Timeout    time.Duration
}

func NewClient(httpClient *fasthttp.Client) *Client {
	return &Client{
		HTTPClient: httpClient,
		Request:    fasthttp.AcquireRequest(),
		Response:   fasthttp.AcquireResponse(),
		Timeout:    5 * time.Second,
	}
}

func (c *Client) SetTimeout(timeout int) {
	c.Timeout = time.Duration(timeout) * time.Second
}

func (c *Client) AddHeader(key, value string) {
	c.Request.Header.Add(key, value)
}

func (c *Client) SetContentType(contentType string) {
	c.Request.Header.SetContentType(contentType)
}

func (c *Client) SetAuthorization(authorization string) {
	c.Request.Header.Set("Authorization", authorization)
}

func (c *Client) executeRequest(method, url string, body []byte) (response []byte, status int) {
	c.Request.SetRequestURI(url)
	c.Request.Header.SetMethod(method)
	c.Request.SetBody(body)

	if err := c.HTTPClient.DoTimeout(c.Request, c.Response, c.Timeout); err != nil {
		fmt.Println("http client error: ", err)
		return nil, fasthttp.StatusInternalServerError
	}
	return c.Response.Body(), c.Response.StatusCode()
}

func (c *Client) Post(url string, body []byte) (response []byte, status int) {
	return c.executeRequest("POST", url, body)
}

func (c *Client) Get(url string) (response []byte, status int) {
	return c.executeRequest("GET", url, nil)
}

func (c *Client) Put(url string, body []byte) (response []byte, status int) {
	return c.executeRequest("PUT", url, body)
}
