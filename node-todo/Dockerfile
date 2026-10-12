# --- Stage 1: Install Dependencies ---
FROM node:20-alpine AS builder
WORKDIR /app
COPY package.json ./   

# PERBAIKAN: Ganti npm ci menjadi npm install
RUN npm install --omit=dev

# --- Stage 2: Runtime Production ---
FROM node:20-alpine
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /app/node_modules ./node_modules
COPY --chown=appuser:appgroup server.js package.json ./

USER appuser
EXPOSE 3000
CMD ["node", "server.js"]
