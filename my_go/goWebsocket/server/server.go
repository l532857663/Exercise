package server

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("vim-go")
	// init
	u := websocket.Upgrader{}
	c, err := u.Upgrade(w, r, nil)
	if err != nil {
		// handle error
	}

	for {
		// receive message
		messageType, message, err := c.ReadMessage()
		if err != nil {
			// handle error
		}

		// send message
		err = c.WriteMessage(messageType, "response")
		if err != nil {
			// handle error
		}
	}
	return
}
