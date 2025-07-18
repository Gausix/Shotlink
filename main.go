package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/emulation"
)

func main() {
	http.HandleFunc("/", handleScreenshot)
	log.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleScreenshot(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "'url' parameter is required", http.StatusBadRequest)
		return
	}

	// Viewport desejado
	const width, height = 1280, 720

	// Create context for chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 15s timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Screenshot in memory
	var buf []byte
	err := chromedp.Run(ctx,
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.CaptureScreenshot(&buf),
	)

	if err != nil {
		http.Error(w, "Error capturing screenshot: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send PNG image as response
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
