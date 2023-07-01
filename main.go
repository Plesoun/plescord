package main

import (
	"flag"
	"log"
	"net/http"

	// local packages
	"main.go/config"
	logger "main.go/logging"

	// external packages
	"github.com/gorilla/websocket"
)

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// Connected clients
var clients = make(map[*websocket.Conn]bool)

// Broadcast channel
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// load configuration
	configPath := flag.String("config", "./localconf.yaml", "Path to configuration file")
	flag.Parse()

	config, configerr := config.LoadConfig(configPath)
	if configerr != nil {
		log.Fatal("Failed to load configuration")
	}

	// initialize logging
	log := logger.NewLogger()
	log.Info("Test")

	// initialize ws
	fs := http.FileServer(http.Dir("./test"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Printf("http server started on :%s", config.WebsocketPort)

	err := http.ListenAndServe(":"+config.WebsocketPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade the the initial GET to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message

		// Read in a new message as JSON and map it to Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error when reading message: %v", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error when handling message: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
