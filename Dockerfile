FROM golang:1.22.4-slim

# Instala dependências + Chrome headless
RUN apt-get update && apt-get upgrade -y && apt-get install -y \
    wget \
    gnupg \
    ca-certificates \
    fonts-liberation \
    libappindicator3-1 \
    libasound2 \
    libatk-bridge2.0-0 \
    libcups2 \
    libgbm1 \
    libgtk-3-0 \
    libnspr4 \
    libnss3 \
    lsb-release \
    xdg-utils \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

# Instala o Chrome
RUN wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list && \
    apt-get update && apt-get install -y google-chrome-stable

# Cria diretório de trabalho
WORKDIR /app

# Copia os arquivos
COPY . .

# Compila o app
RUN go mod tidy
RUN go build -o server .

# Expõe a porta do servidor
EXPOSE 8080

# Executa o binário
CMD ["./server"]
