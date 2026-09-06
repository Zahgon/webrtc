//go:build !js

package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/pion/webrtc/v4"
)

var (
	tracksLock sync.RWMutex
	tracks     []*webrtc.TrackLocalStaticSample

	videoFiles     []string
	videoFileIndex atomic.Int32
)

func nextVideo() { _ = "STUB: not implemented"; return }

func doWHIP(res http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func playFile(fileIndex int32) { _ = "STUB: not implemented"; return }

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/whip", doWHIP)

	go func() {
		files, err := filepath.Glob("*.ivf")
		if err != nil {
			panic(err)
		}
		for _, p := range files {
			videoFiles = append(videoFiles, filepath.Base(p))
		}
		if len(videoFiles) == 0 {
			panic("no .ivf files found in the working directory")
		}

		for {
			playFile(videoFileIndex.Load())
		}
	}()

	fmt.Println("Open http://localhost:8080 to access this demo")

	panic(http.ListenAndServe(":8080", nil))
}
