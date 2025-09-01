FROM golang:1.24.5-bookworm

# Instala dependências do sistema e do Chrome
RUN apt-get update && apt-get install -y \
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
    xdg-utils && \
    rm -rf /var/lib/apt/lists/*

# Instala o Google Chrome
RUN wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list && \
    apt-get update && apt-get install -y google-chrome-stable && \
    rm -rf /var/lib/apt/lists/*

# Define diretório do app
WORKDIR /app

# Copia o projeto
COPY . .

# Resolve dependências
RUN go mod tidy

# Compila o binário
RUN go build -o server .

# Expõe a porta
EXPOSE 8080

# Executa
CMD ["./server"]
