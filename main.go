package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

func main() {
	http.HandleFunc("/get", handleScreenshot)
	log.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleScreenshot(writer http.ResponseWriter, r *http.Request) {
	// Get URL parameter
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(writer, "'url' parameter is required", http.StatusBadRequest)
		return
	}

	// Get width and height from query parameters, or use defaults
	width := 1280
	height := 720

	if widthStr := r.URL.Query().Get("width"); widthStr != "" {
		if parsedWidth, err := strconv.Atoi(widthStr); err == nil && parsedWidth > 0 {
			width = parsedWidth
		} else {
			http.Error(writer, "Invalid width parameter", http.StatusBadRequest)
			return
		}
	}

	if heightStr := r.URL.Query().Get("height"); heightStr != "" {
		if parsedHeight, err := strconv.Atoi(heightStr); err == nil && parsedHeight > 0 {
			height = parsedHeight
		} else {
			http.Error(writer, "Invalid height parameter", http.StatusBadRequest)
			return
		}
	}

	log.Printf("Taking screenshot with dimensions: %dx%d", width, height)

	// Create context for chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 15s timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Screenshot in memory
	var buf []byte
	err := chromedp.Run(ctx,
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.Sleep(2*time.Second), // Wait for page to load
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.CaptureScreenshot(&buf).Do(ctx)
		}),
	)

	if err != nil {
		http.Error(writer, "Error capturing screenshot: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send PNG image as response
	writer.Header().Set("Content-Type", "image/png")
	writer.WriteHeader(http.StatusOK)
	writer.Write(buf)
}
