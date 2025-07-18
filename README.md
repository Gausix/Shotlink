# Snaptor
Website screenshot API service

# How to use

1. **Make a request**: Send a GET request to `https://snaptor.httpshield.net/get?url=<URL>`, replacing `<URL>` with the website you want to capture.
2. **Receive the screenshot**: The server will respond with a PNG image of the website screenshot.

# Example request
```bash
curl -X GET "https://snaptor.httpshield.net/get?url=https://example.com" -o screenshot.png
```

# Requirements
- Go 1.18 or later
- `chromedp` package for browser automation
- `cdproto` package for Chrome DevTools Protocol
- `emulation` package for viewport emulation
- `context` package for managing request context
- `time` package for request timeout
- `net/http` package for handling HTTP requests
- `log` package for logging errors

# How to run locally

1. Clone the repository:
   ```bash
   git clone https://github.com/Gausix/Snaptor
   ```
2. Change into the project directory:
   ```bash
   cd Snaptor
   ```
3. Install the required Go packages:
   ```bash
   go mod tidy
   ```
4. Run the server:
   ```bash
   go run main.go
   ```
5. Access the service at `http://localhost:8080/get?url=<URL>`. Replace `<URL>` with the website you want to capture.