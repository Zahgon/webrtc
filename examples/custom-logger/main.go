//go:build !js

package main

import (
	"fmt"
	"os"

	"github.com/pion/logging"
	"github.com/pion/webrtc/v4"
)

type customLogger struct{}

func (c customLogger) Trace(string)          { _ = "STUB: not implemented"; return }
func (c customLogger) Tracef(string, ...any) { _ = "STUB: not implemented"; return }

func (c customLogger) Debug(msg string)                  { _ = "STUB: not implemented"; return }
func (c customLogger) Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (c customLogger) Info(msg string)                  { _ = "STUB: not implemented"; return }
func (c customLogger) Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

func (c customLogger) Warn(msg string)                  { _ = "STUB: not implemented"; return }
func (c customLogger) Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (c customLogger) Error(msg string)                  { _ = "STUB: not implemented"; return }
func (c customLogger) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

type customLoggerFactory struct{}

func (c customLoggerFactory) NewLogger(subsystem string) logging.LeveledLogger {
	_ = "STUB: not implemented"
	return *new(logging.LeveledLogger)
}

func main() {

	s := webrtc.SettingEngine{
		LoggerFactory: customLoggerFactory{},
	}
	api := webrtc.NewAPI(webrtc.WithSettingEngine(s))

	offerPeerConnection, err := api.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		panic(err)
	}
	defer func() {
		if cErr := offerPeerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close offerPeerConnection: %v\n", cErr)
		}
	}()

	if _, err = offerPeerConnection.CreateDataChannel("custom-logger", nil); err != nil {
		panic(err)
	}

	answerPeerConnection, err := api.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		panic(err)
	}
	defer func() {
		if cErr := answerPeerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close answerPeerConnection: %v\n", cErr)
		}
	}()

	offerPeerConnection.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s (offerer)\n", state.String())

		if state == webrtc.PeerConnectionStateFailed {

			fmt.Println("Peer Connection has gone to failed exiting")
			os.Exit(0)
		}

		if state == webrtc.PeerConnectionStateClosed {

			fmt.Println("Peer Connection has gone to closed exiting")
			os.Exit(0)
		}
	})

	answerPeerConnection.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s (answerer)\n", state.String())

		if state == webrtc.PeerConnectionStateFailed {

			fmt.Println("Peer Connection has gone to failed exiting")
			os.Exit(0)
		}
	})

	answerPeerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			if iceErr := offerPeerConnection.AddICECandidate(candidate.ToJSON()); iceErr != nil {
				panic(iceErr)
			}
		}
	})

	offerPeerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			if iceErr := answerPeerConnection.AddICECandidate(candidate.ToJSON()); iceErr != nil {
				panic(iceErr)
			}
		}
	})

	offer, err := offerPeerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	if err = offerPeerConnection.SetLocalDescription(offer); err != nil {
		panic(err)
	}

	if err = answerPeerConnection.SetRemoteDescription(offer); err != nil {
		panic(err)
	}

	answer, err := answerPeerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	if err = answerPeerConnection.SetLocalDescription(answer); err != nil {
		panic(err)
	}

	if err = offerPeerConnection.SetRemoteDescription(answer); err != nil {
		panic(err)
	}

	select {}
}
