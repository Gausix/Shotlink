<div align='center'>
   <img src='https://i.imgur.com/yejUTRS.png' width='96px' style='text-align:center;' /> 
</div>

[![Deploy on Railway](https://railway.com/button.svg)](https://railway.com/deploy/shotlink?referralCode=U520U6)

**Shotlink** is a lightweight and efficient REST API service written in Go that allows you to capture high-quality screenshots of websites programmatically. Designed for developers, automation tools, monitoring systems, and visual reporting platforms, Shotlink leverages Chrome Headless via `chromedp` to render fully interactive web pages—including dynamic JavaScript content and custom styles.

# How to use

1. **Make a request**: Send a GET request to `https://snaptor.nexvul.com/get?url=<URL>`, replacing `<URL>` with the website you want to capture.
2. **Receive the screenshot**: The server will respond with a PNG image of the website screenshot.

# Example request

```bash
curl -X GET "https://snaptor.nexvul.com/get?url=https://example.com" -o screenshot.png
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
