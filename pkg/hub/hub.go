package hub

import (
	"encoding/json"
	"log"
	"net"

	"github.com/ping-pong/pkg/models"
)

// Hub holds data structure for persisting references to incoming connections
// and channels to communicate through
type Hub struct {
	clients    map[int]*models.HubClient
	register   chan *models.HubClient
	unregister chan *models.HubClient
	shutDown   chan bool
	listener   net.Listener
}

// Start spins up the hub and listens for incoming connections
func Start() (*Hub, error) {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		return nil, err
	}

	hub := &Hub{
		clients:    make(map[int]*models.HubClient),
		register:   make(chan *models.HubClient),
		unregister: make(chan *models.HubClient),
		shutDown:   make(chan bool),
		listener:   listener,
	}

	go hub.listen()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil || conn == nil {
				return
			}

			c, err := hub.getClient(conn)
			if err != nil {
				conn.Close()
				continue
			}

			// make sure user with same id as existing online user doesn't get in
			if _, ok := hub.clients[c.ID]; ok {
				conn.Write([]byte("user already connected"))
				continue
			}

			hub.register <- c
			go hub.receive(c)
			go hub.transmit(c)
		}
	}()

	return hub, nil
}

func (h *Hub) getClient(conn net.Conn) (*models.HubClient, error) {
	msgBuff := make([]byte, 4096)
	msgLen, err := conn.Read(msgBuff)
	if err != nil {
		return nil, err
	}

	msg := &models.Message{}
	if err := json.Unmarshal(msgBuff[:msgLen], msg); err != nil {
		return nil, err
	}

	return &models.HubClient{
		Socket:  conn,
		Data:    make(chan []byte),
		ID:      msg.UserID,
		Friends: msg.Friends,
	}, nil
}

// spins up channel consumer to keep track of clinet join/leave and application exit
func (h *Hub) listen() {
	for {
		select {
		case c := <-h.register:
			h.updateStatus(c, []byte(`{"online":true}`))
			h.clients[c.ID] = c
			log.Printf("user ID:%d has joined", c.ID)
		case c := <-h.unregister:
			if _, ok := h.clients[c.ID]; ok {
				h.updateStatus(c, []byte(`{"online":false}`))
				close(c.Data)
				delete(h.clients, c.ID)
				log.Printf("user ID:%d has left", c.ID)
			}
		case <-h.shutDown:
			return
		}
	}
}

// UpdateStatus writes message to client with online true/false flag
func (h *Hub) updateStatus(hc *models.HubClient, msg []byte) {
	for _, id := range hc.Friends {
		if c, ok := h.clients[id]; ok {
			c.Data <- msg
		}
	}
}

func (h *Hub) receive(hc *models.HubClient) {
	defer hc.Socket.Close()
	for {
		message := make([]byte, 4096)
		length, err := hc.Socket.Read(message)

		if err != nil {
			h.unregister <- hc
			return
		}

		if length > 0 {
			hc.Data <- message
		}
	}
}

func (h *Hub) transmit(hc *models.HubClient) {
	defer hc.Socket.Close()
	for message := range hc.Data {
		if _, err := hc.Socket.Write(message); err != nil {
			h.unregister <- hc
			return
		}
	}
}

// Shutdown notifies all clients about shutdown, closes clients and hub
func (h *Hub) Shutdown() {
	// stop server
	h.listener.Close()

	// stop channels from listening for messages
	h.shutDown <- true

	// close all client connections
	for _, c := range h.clients {
		c.Socket.Write([]byte("server shutdown"))
		c.Socket.Close()
	}
}
