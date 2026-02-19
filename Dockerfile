# ── Stage 1: Build ──────────────────────────────────────────────
FROM golang:1.26-alpine AS builder

WORKDIR /src

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build a statically-linked binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -trimpath -o /bin/server ./cmd

# ── Stage 2: Runtime ───────────────────────────────────────────
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /bin/server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]
