//go:build !js

package main

import (
	"context"
	"embed"
	"flag"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media/oggreader"
)

const (
	playlistFile = "playlist.ogg"
	labelAudio   = "audio"
	labelTrack   = "pion"
)

//go:embed web/*
var content embed.FS

type bufferedPage struct {
	payload  []byte
	duration time.Duration
	granule  uint64
}

type oggTrack struct {
	serial uint32
	header *oggreader.OggHeader
	tags   *oggreader.OpusTags

	title   string
	artist  string
	vendor  string
	pages   []bufferedPage
	runtime time.Duration
}

func main() { //nolint:gocognit,cyclop
	addr := flag.String("addr", "localhost:8080", "HTTP listen address")
	flag.Parse()

	tracks, err := parsePlaylist(playlistFile)
	if err != nil {
		log.Fatal(err)
	}
	if len(tracks) == 0 {
		log.Fatal("no playable Opus pages were found in playlist.ogg")
	}

	log.Printf("Loaded %d track(s) from %s", len(tracks), playlistFile)
	for i, t := range tracks {
		log.Printf("  [%d] serial=%d title=%q artist=%q pages=%d duration=%v",
			i+1, t.serial, t.title, t.artist, len(t.pages), t.runtime)
	}

	static, err := fs.Sub(content, "web")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.FS(static))
	mux.Handle("/", fileServer)
	mux.HandleFunc("/whep", func(writer http.ResponseWriter, reader *http.Request) {
		if reader.Method != http.MethodPost {
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)

			return
		}

		body, err := io.ReadAll(reader.Body)
		if err != nil {
			http.Error(writer, "failed to read body", http.StatusBadRequest)

			return
		}
		rawSDP := string(body)
		if strings.TrimSpace(rawSDP) == "" {
			http.Error(writer, "empty SDP", http.StatusBadRequest)

			return
		}
		log.Printf("received offer (%d bytes)", len(rawSDP))

		offer := webrtc.SessionDescription{Type: webrtc.SDPTypeOffer, SDP: rawSDP}

		answer, err := handleOffer(tracks, offer) //nolint:contextcheck
		if err != nil {
			log.Printf("error handling offer: %v", err)
			http.Error(writer, err.Error(), http.StatusBadRequest)

			return
		}

		writer.Header().Set("Content-Type", "application/sdp")
		if _, err = writer.Write([]byte(answer.SDP)); err != nil {
			log.Printf("write answer failed: %v", err)
		}
	})

	log.Printf("Serving UI at http://%s ...", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux)) //nolint:gosec
}

//nolint:cyclop
func handleOffer(
	tracks []*oggTrack,
	offer webrtc.SessionDescription,
) (*webrtc.SessionDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

//nolint:contextcheck // webrtc API does not take context for SetRemoteDescription

func stream(
	tracks []*oggTrack,
	audioTrack *webrtc.TrackLocalStaticSample,
	currentTrack *atomic.Int32,
	switchTrack <-chan int,
	playlistChannel *webrtc.DataChannel,
	ctx context.Context,
) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

func parsePlaylist(path string) ([]*oggTrack, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil, nil
}

//nolint:err113

//nolint:gosec // path is validated and confined to local directory

//nolint:nestif

func ensureTrack(tracks map[uint32]*oggTrack, serial uint32, order *[]uint32) *oggTrack {
	_ = "STUB: not implemented"
	return nil
}

func extractMetadata(tags *oggreader.OpusTags) (title, artist string) {
	_ = "STUB: not implemented"
	return "", ""
}

func pageDuration(header *oggreader.OggHeader, granule, last uint64) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

//nolint:gosec

func wrapNext(current, limit int) int { _ = "STUB: not implemented"; return 0 }

func wrapPrev(current, limit int) int { _ = "STUB: not implemented"; return 0 }

func normalizeIndex(i, limit int) int { _ = "STUB: not implemented"; return 0 }

func sendPlaylistText(dc *webrtc.DataChannel, tracks []*oggTrack, current int, includeNow bool) {
	_ = "STUB: not implemented"
	return
}

func sendNowPlayingText(dc *webrtc.DataChannel, track *oggTrack, index int) {
	_ = "STUB: not implemented"
	return
}

func nowLine(track *oggTrack, index int) string { _ = "STUB: not implemented"; return "" }

func cleanText(v string) string { _ = "STUB: not implemented"; return "" }
