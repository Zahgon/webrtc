//go:build !js

package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/pion/webrtc/v4"
)

var (
	peerConnectionConfiguration = webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	broadcastHub = &Hub{
		connections: make(map[*webrtc.DataChannel]bool),
		usernames:   make(map[*webrtc.DataChannel]string),
		mu:          sync.RWMutex{},
	}
)

type Hub struct {
	connections map[*webrtc.DataChannel]bool
	usernames   map[*webrtc.DataChannel]string
	mu          sync.RWMutex
}

var (
	adjectives = []string{
		"Quick", "Swift", "Bright", "Bold", "Calm", "Cool", "Fast", "Happy",
		"Lucky", "Shy", "Sneaky", "Wise", "Brave", "Clever", "Kind", "Proud",
	}
	nouns = []string{
		"Fox", "Eagle", "Lion", "Tiger", "Wolf", "Dragon", "Hawk", "Bear",
		"Shark", "Falcon", "Leopard", "Panther", "Phoenix", "Raven", "Crow", "Owl",
	}
)

func (h *Hub) Register(channel *webrtc.DataChannel) string { _ = "STUB: not implemented"; return "" }

func (h *Hub) Unregister(channel *webrtc.DataChannel) { _ = "STUB: not implemented"; return }

func (h *Hub) generateUniqueUsername() string { _ = "STUB: not implemented"; return "" }

func (h *Hub) GetUsername(channel *webrtc.DataChannel) string { _ = "STUB: not implemented"; return "" }

func (h *Hub) Broadcast(message string, sender *webrtc.DataChannel) {
	_ = "STUB: not implemented"
	return
}

func (h *Hub) Count() int { _ = "STUB: not implemented"; return 0 }

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/whep", whepHandler)
	http.HandleFunc("/whip", whipHandler)

	fmt.Println("Open http://localhost:8080 to access this demo")
	panic(http.ListenAndServe(":8080", nil))
}

func whipHandler(res http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func whepHandler(res http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func writeAnswer(res http.ResponseWriter, peerConnection *webrtc.PeerConnection, offer []byte, path string) {
	_ = "STUB: not implemented"
	return
}

//nolint: errcheck
