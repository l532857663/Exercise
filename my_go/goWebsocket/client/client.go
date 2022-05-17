package client

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct {
	Url  *url.URL
	Conn *websocket.Conn
}

func NewClient(host, path, scheme string) (*Client, error) {
	// schema – can be ws:// or wss://
	// host, port – WebSocket server
	if path == "" {
		path = "/"
	}
	u := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		// handle error
		return nil, err
	}

	client := &Client{
		Url:  u,
		Conn: conn,
	}

	return client, nil
}

func (c *Client) TheSocketConnInfo() {
	fmt.Printf("The host is: [%s]\n", c.Url.Host)
	return
}

func (c *Client) Close() {
	c.Conn.Close()
	return
}

func (c *Client) SendMessage(message []byte) []byte {
	conn := c.Conn
	// send message
	err := conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		// handle error
		return nil
	}
	// receive message
	_, message, err = conn.ReadMessage()
	if err != nil {
		// handle error
		return nil
	}
	return message
}
