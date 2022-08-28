package main

import "fmt"

type Client struct {
	id string
}

func NewClient(id string) *Client {
	return &Client{
		id: id,
	}
}

func (c *Client) Update(placeName string) {
	fmt.Printf("Dear %s, motion detected in %s!\n", c.id, placeName)
}

func (c *Client) GetId() string {
	return c.id
}
