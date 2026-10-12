# --- Stage 1: Build Binary ---
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go-notes main.go

# --- Stage 2: Runtime Production ---
FROM alpine:3.19
# Membuat non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /app
COPY --from=builder /app/go-notes .

EXPOSE 8080
CMD ["./go-notes"]