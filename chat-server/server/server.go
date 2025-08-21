package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

// Client represents a connected client
type Client struct {
	conn net.Conn
	name string
	room string // New field: current room name
}

// ClientMessage represents a message from a client
type ClientMessage struct {
	Client  *Client
	Message string
}

// Server represents the chat server
type Server struct {
	clients    map[net.Conn]*Client
	usernames  map[string]*Client
	rooms      map[string]map[string]*Client // New field: roomName -> clientName -> Client
	messages   chan ClientMessage
	register   chan *Client
	unregister chan *Client
	mutex      sync.Mutex
}

// NewServer creates a new chat server
func NewServer() *Server {
	return &Server{
		clients:    make(map[net.Conn]*Client),
		usernames:  make(map[string]*Client),
		rooms:      make(map[string]map[string]*Client), // Initialize rooms map
		messages:   make(chan ClientMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Start starts the chat server
func (s *Server) Start(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()
	log.Printf("Chat server started on port %s", port)

	go s.handleMessages()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleMessages() {
	for {
		select {
		case client := <-s.register:
			s.mutex.Lock()
			s.clients[client.conn] = client
			s.usernames[client.name] = client
			// Add client to their initial room
			if _, ok := s.rooms[client.room]; !ok {
				s.rooms[client.room] = make(map[string]*Client)
			}
			s.rooms[client.room][client.name] = client
			s.mutex.Unlock()
			s.broadcastMessageToRoom(client.room, fmt.Sprintf("%s has joined the chat.", client.name))
			log.Printf("Client %s connected to room %s. Total clients: %d", client.name, client.room, len(s.clients))

		case client := <-s.unregister:
			s.mutex.Lock()
			if _, ok := s.clients[client.conn]; ok {
				delete(s.clients, client.conn)
				delete(s.usernames, client.name)
				// Remove client from their room
				if roomClients, ok := s.rooms[client.room]; ok {
					delete(roomClients, client.name)
					if len(roomClients) == 0 { // If room is empty, delete it
						delete(s.rooms, client.room)
					}
				}
				client.conn.Close()
				s.mutex.Unlock()
				s.broadcastMessageToRoom(client.room, fmt.Sprintf("%s has left the chat.", client.name))
				log.Printf("Client %s disconnected from room %s. Total clients: %d", client.name, client.room, len(s.clients))
			} else {
				s.mutex.Unlock()
			}

		case clientMsg := <-s.messages:
			senderClient := clientMsg.Client
			actualMessage := clientMsg.Message

			if strings.HasPrefix(actualMessage, "/whisper ") {
				whisperParts := strings.SplitN(actualMessage, " ", 3)
				if len(whisperParts) < 3 {
					fmt.Fprintln(senderClient.conn, "Usage: /whisper <username> <message>")
					continue
				}
				targetUsername := whisperParts[1]
				privateMessage := whisperParts[2]

				s.sendWhisper(senderClient, targetUsername, privateMessage)
			} else if strings.HasPrefix(actualMessage, "/join ") {
				newRoomName := strings.TrimSpace(strings.TrimPrefix(actualMessage, "/join "))
				s.joinRoom(senderClient, newRoomName)
			} else if actualMessage == "/leave" {
				s.leaveRoom(senderClient)
			} else {
				// Broadcast to current room
				s.broadcastMessageToRoom(senderClient.room, fmt.Sprintf("%s: %s", senderClient.name, actualMessage))
			}
		}
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	fmt.Fprint(conn, "Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading name: %v", err)
		conn.Close()
		return
	}
	name = strings.TrimSpace(name) // Remove newline and any other whitespace

	if name == "" {
		fmt.Fprintln(conn, "Name cannot be empty. Disconnecting.")
		conn.Close()
		return
	}

	// Check if name is already taken
	s.mutex.Lock()
	if _, exists := s.usernames[name]; exists {
		s.mutex.Unlock()
		fmt.Fprintln(conn, fmt.Sprintf("Name '%s' is already taken. Please choose another. Disconnecting.", name))
		conn.Close()
		return
	}
	s.mutex.Unlock()

	client := &Client{conn: conn, name: name, room: "general"} // Assign default room
	s.register <- client

	defer func() {
		s.unregister <- client
	}()

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading from %s: %v", client.name, err)
			break
		}
		s.messages <- ClientMessage{Client: client, Message: strings.TrimSpace(message)}
	}
}

// broadcastMessageToRoom sends a message to all clients in a specific room.
func (s *Server) broadcastMessageToRoom(roomName string, message string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if roomClients, ok := s.rooms[roomName]; ok {
		for _, client := range roomClients {
			_, err := fmt.Fprintln(client.conn, message)
			if err != nil {
				log.Printf("Error sending message to %s in room %s: %v", client.name, roomName, err)
			}
		}
	}
}

// sendWhisper sends a private message from senderClient to targetUsername.
func (s *Server) sendWhisper(senderClient *Client, targetUsername, msg string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if senderClient.name == targetUsername {
		fmt.Fprintln(senderClient.conn, "You cannot whisper to yourself.")
		return
	}

	targetClient, found := s.usernames[targetUsername]
	if !found {
		fmt.Fprintln(senderClient.conn, fmt.Sprintf("User '%s' not found.", targetUsername))
		return
	}

	// Send to target
	_, err := fmt.Fprintln(targetClient.conn, fmt.Sprintf("[Whisper from %s]: %s", senderClient.name, msg))
	if err != nil {
		log.Printf("Error sending whisper to %s: %v", targetUsername, err)
	}

	// Send confirmation to sender
	_, err = fmt.Fprintln(senderClient.conn, fmt.Sprintf("[Whisper to %s]: %s", targetUsername, msg))
	if err != nil {
		log.Printf("Error sending whisper confirmation to %s: %v", senderClient.name, err)
	}
}

// joinRoom handles a client joining a new room.
func (s *Server) joinRoom(client *Client, newRoomName string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if newRoomName == "" {
		fmt.Fprintln(client.conn, "Room name cannot be empty.")
		return
	}

	if client.room == newRoomName {
		fmt.Fprintln(client.conn, fmt.Sprintf("You are already in room '%s'.", newRoomName))
		return
	}

	// Remove client from old room
	if oldRoomClients, ok := s.rooms[client.room]; ok {
		delete(oldRoomClients, client.name)
		if len(oldRoomClients) == 0 {
			delete(s.rooms, client.room) // Delete room if empty
		}
		s.broadcastMessageToRoom(client.room, fmt.Sprintf("%s has left the room.", client.name))
	}

	// Add client to new room
	if _, ok := s.rooms[newRoomName]; !ok {
		s.rooms[newRoomName] = make(map[string]*Client)
	}
	s.rooms[newRoomName][client.name] = client
	client.room = newRoomName // Update client's room

	fmt.Fprintln(client.conn, fmt.Sprintf("You have joined room '%s'.", newRoomName))
	s.broadcastMessageToRoom(newRoomName, fmt.Sprintf("%s has joined the room.", client.name))
	log.Printf("Client %s joined room %s.", client.name, newRoomName)
}

// leaveRoom handles a client leaving their current room.
func (s *Server) leaveRoom(client *Client) {
	// Simply call joinRoom to move to general
	s.joinRoom(client, "general")
}
