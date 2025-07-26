FROM golang:1.24.5-alpine@sha256:daae04ebad0c21149979cd8e9db38f565ecefd8547cf4a591240dc1972cf1399 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy source code
COPY . .

RUN go build -o app .

FROM alpine:latest@sha256:4bcff63911fcb4448bd4fdacec207030997caf25e9bea4045fa6c8c44de311d1

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE ${API_PORT:-8080}

CMD ["./app"]
