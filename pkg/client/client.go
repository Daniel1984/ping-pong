package client

import (
	"fmt"
	"log"
	"net"
	"os"
)

// Client holds reference to network connection
type Client struct {
	socket net.Conn
}

// Start initiates TCP connection to server, starts listening for incoming
// messages and returns instance of the client and/or error
func Start(payload []byte) (*Client, error) {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		return nil, err
	}

	c := &Client{socket: conn}
	go c.receive()

	if _, err := conn.Write(payload); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) receive() {
	for {
		msg := make([]byte, 4096)
		length, err := c.socket.Read(msg)
		if err != nil {
			c.socket.Close()
			return
		}

		if length > 0 {
			switch v := fmt.Sprintf("%s", msg[:length]); {
			case v == "user already connected":
				log.Printf("%s", v)
				c.socket.Close()
				os.Exit(1)
			case v == "server shutdown":
				log.Printf("%s", v)
				c.socket.Close()
				os.Exit(1)
			default:
				log.Printf("%s", v)
			}
		}
	}
}

// Close terminates socket connection
func (c *Client) Close() {
	c.socket.Close()
}
