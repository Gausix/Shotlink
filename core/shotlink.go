package core

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

var (
	parentCtx context.Context
	cancelCtx context.CancelFunc
)

func init() {
	parentCtx, cancelCtx = chromedp.NewContext(context.Background())
}

func Shutdown() {
	if cancelCtx != nil {
		cancelCtx()
	}
}

func HandleScreenshot(writer http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(writer, "'url' parameter is required", http.StatusBadRequest)
		return
	}

	width, height := 1280, 720
	if w := r.URL.Query().Get("w"); w != "" {
		if val, err := strconv.Atoi(w); err == nil {
			width = val
		}
	}
	if h := r.URL.Query().Get("h"); h != "" {
		if val, err := strconv.Atoi(h); err == nil {
			height = val
		}
	}

	timeout := 30 * time.Second
	if t := os.Getenv("SCREENSHOT_TIMEOUT"); t != "" {
		if val, err := strconv.Atoi(t); err == nil {
			timeout = time.Duration(val) * time.Second
		}
	}

	log.Printf("Screenshot -> url=%s, size=%dx%d", url, width, height)

	ctx, cancel := chromedp.NewContext(parentCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.FullScreenshot(&buf, 90),
	)

	if err != nil {
		log.Printf("Error capturing screenshot (%s): %v", url, err)
		http.Error(writer, "Error capturing screenshot: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "image/png")
	writer.Header().Set("Cache-Control", "no-store")
	writer.WriteHeader(http.StatusOK)
	writer.Write(buf)
}
