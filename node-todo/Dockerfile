# --- Stage 1: Install Dependencies ---
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install --omit=dev

# --- Stage 2: Runtime Production ---
# Gunakan tag alpine terbaru untuk mendapatkan versi OpenSSL yang aman
FROM node:20-alpine  

# Solusi mitigasi tambahan: Paksa apk memperbarui libcrypto & libssl ke versi terbaru
RUN apk update && apk add --no-cache libcrypto3 libssl3

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app

COPY --from=builder /app/node_modules ./node_modules
COPY --chown=appuser:appgroup server.js package*.json ./

USER appuser
EXPOSE 3000
CMD ["node", "server.js"]
