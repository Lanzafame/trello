// Copyright 2016 The trello client Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trello

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	trelloEndpoint = "https://api.trello.com/1"
	jbt            = "application/json; charset=utf-8"
)

// A Client represents a Trello client.
type Client struct {
	key   string
	token string

	*http.Client
}

// NewClient returns a Trello client.
func NewClient(key, token string) *Client {
	c := &Client{
		key:   key,
		token: token,
		&http.Client{},
	}
	return c
}

func (c *Client) Get(url string) (*http.Response, error) {
	url = c.mustAppendKeyToken(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) Post(url string, body io.Reader) (*http.Response, error) {
	url = c.mustAppendKeyToken(url)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", jbt)
	return c.Do(req)
}

func (c *Client) Put(url string, body io.Reader) (*http.Response, error) {
	url = c.mustAppendKeyToken(url)
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", jbt)
	return c.Do(req)
}

func (c *Client) Delete(url string, body io.Reader) (*http.Response, error) {
	url = c.mustAppendKeyToken(url)
	req, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", jbt)
	return c.Do(req)
}

func (c *Client) appendKeyToken(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	params := url.Query()
	params.Add("key", c.key)
	params.Add("token", c.token)
	url.RawQuery = params.Encode()

	return url.String(), nil
}

func (c *Client) mustAppendKeyToken(rawurl string) string {
	url, err := c.appendKeyToken(rawurl)
	if err != nil {
		panic(err)
	}
	return url
}

func (c *Client) Boards(username string) ([]Board, error) {
	url := fmt.Sprintf("%v/members/%v/boards", trelloEndpoint, username)
	resp, err := c.boardEndpoint(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	boards := []Board{}
	if err = json.Unmarshal(resp.Body, &boards); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code (%v)", resp.StatusCode)
	}

	return boards, nil
}

func (c *Client) boardEndpoint(url, httpMethod string, body io.Reader) (*http.Response, error) {
	var err error
	var resp *http.Response
	switch httpMethod {
	case "GET":
		resp, err = c.Get(url)
	case "POST":
		resp, err = c.Post(url, body)
	case "PUT":
		resp, err = c.Put(url, body)
	case "DELETE":
		resp, err = c.Delete(url, body)
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Lists returns the lists under the given board.
func (c *Client) Lists(boardID string) ([]List, error) {
	url := fmt.Sprintf("%v/boards/%v/lists", trelloEndpoint, boardID)
	body, err := c.Get(url)
	if err != nil {
		return nil, err
	}

	lists := []List{}
	if err := json.Unmarshal(body, &lists); err != nil {
		return nil, err
	}

	return lists, nil
}

// Labels returns the labels under the given board.
func (c *Client) Labels(boardID string) ([]Label, error) {
	url := fmt.Sprintf("%v/boards/%v/labels", trelloEndpoint, boardID)
	body, err := c.Get(url)
	if err != nil {
		return nil, err
	}

	labels := []Label{}
	if err := json.Unmarshal(body, &labels); err != nil {
		return nil, err
	}

	return labels, nil
}

// PushCard creates a new card in trello.
func (c *Client) PushCard(card Card) error {
	url := fmt.Sprintf("%v/cards", trelloEndpoint)

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(card); err != nil {
		return err
	}

	if _, err := c.Post(url, buf); err != nil {
		return err
	}

	return nil
}
