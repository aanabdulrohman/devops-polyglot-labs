FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENAMBLE=0 GOOS=linux go build -o main.go

FROM alpine:3.19
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /app
COPY --from=builder /app/go-notes .

EXPOSE 8080
CMD ["./go-notes"]