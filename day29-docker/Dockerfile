# Stage 1 - Builder
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# Cache deps before copying source
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./main.go

# Stage 2 - Final Image (~15MB)
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

# Non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/.env.example .

RUN chown -R appuser:appgroup /app
USER appuser

EXPOSE 8080

CMD ["./server"]
