//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/pion/webrtc/v4"
)

func main() {

	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	pc, err := webrtc.NewPeerConnection(config)
	if err != nil {
		handleError(err)
	}

	sendChannel, err := pc.CreateDataChannel("foo", nil)
	if err != nil {
		handleError(err)
	}
	sendChannel.OnClose(func() {
		fmt.Println("sendChannel has closed")
	})
	sendChannel.OnClosing(func() {
		fmt.Println("sendChannel is closing")
	})
	sendChannel.OnError(func(err error) {
		fmt.Println("sendChannel error", err)
	})
	sendChannel.OnOpen(func() {
		fmt.Println("sendChannel has opened")

		candidatePair, err := pc.SCTP().Transport().ICETransport().GetSelectedCandidatePair()

		fmt.Println(candidatePair)
		fmt.Println(err)
	})
	sendChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		log(fmt.Sprintf("Message from DataChannel %s payload %s", sendChannel.Label(), string(msg.Data)))
	})

	offer, err := pc.CreateOffer(nil)
	if err != nil {
		handleError(err)
	}
	if err := pc.SetLocalDescription(offer); err != nil {
		handleError(err)
	}

	pc.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		log(fmt.Sprint(state))
	})
	pc.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			encodedDescr := encode(pc.LocalDescription())
			el := getElementByID("localSessionDescription")
			el.Set("value", encodedDescr)
		}
	})

	js.Global().Set("sendMessage", js.FuncOf(func(_ js.Value, _ []js.Value) any {
		go func() {
			el := getElementByID("message")
			message := el.Get("value").String()
			if message == "" {
				js.Global().Call("alert", "Message must not be empty")
				return
			}
			if err := sendChannel.SendText(message); err != nil {
				handleError(err)
			}
		}()
		return js.Undefined()
	}))
	js.Global().Set("startSession", js.FuncOf(func(_ js.Value, _ []js.Value) any {
		go func() {
			el := getElementByID("remoteSessionDescription")
			sd := el.Get("value").String()
			if sd == "" {
				js.Global().Call("alert", "Session Description must not be empty")
				return
			}

			descr := webrtc.SessionDescription{}
			decode(sd, &descr)
			if err := pc.SetRemoteDescription(descr); err != nil {
				handleError(err)
			}
		}()
		return js.Undefined()
	}))
	js.Global().Set("copySDP", js.FuncOf(func(_ js.Value, _ []js.Value) any {
		go func() {
			defer func() {
				if e := recover(); e != nil {
					switch e := e.(type) {
					case error:
						handleError(e)
					default:
						handleError(fmt.Errorf("recovered with non-error value: (%T) %s", e, e))
					}
				}
			}()

			browserSDP := getElementByID("localSessionDescription")

			browserSDP.Call("focus")
			browserSDP.Call("select")

			copyStatus := js.Global().Get("document").Call("execCommand", "copy")
			if copyStatus.Bool() {
				log("Copying SDP was successful")
			} else {
				log("Copying SDP was unsuccessful")
			}
		}()
		return js.Undefined()
	}))

	select {}
}

func log(msg string) { _ = "STUB: not implemented"; return }

func handleError(err error) { _ = "STUB: not implemented"; return }

func getElementByID(id string) js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func readUntilNewline() (in string) { _ = "STUB: not implemented"; return "" }

func encode(obj *webrtc.SessionDescription) string { _ = "STUB: not implemented"; return "" }

func decode(in string, obj *webrtc.SessionDescription) { _ = "STUB: not implemented"; return }
