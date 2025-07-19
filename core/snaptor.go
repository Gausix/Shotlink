package core

import (
	"log"
	"time"
	"context"
	"net/http"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/emulation"
)

func HandleScreenshot(writer http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(writer, "'url' parameter is required", http.StatusBadRequest)
		return
	}

	width, height := 1280, 720
	log.Printf("Taking screenshot with dimensions: %dx%d", width, height)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.Sleep(2*time.Second),
		
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.CaptureScreenshot(&buf).Do(ctx)
		}),
	)

	if err != nil {
		http.Error(writer, "Error capturing screenshot: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "image/png")
	writer.WriteHeader(http.StatusOK)
	writer.Write(buf)
}
