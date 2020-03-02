package models

import "net"

// HubClient structure to hold calling client reference
type HubClient struct {
	Socket  net.Conn
	Data    chan []byte
	ID      int
	Friends []int
}
