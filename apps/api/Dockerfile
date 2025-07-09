FROM golang:1.24.5-alpine@sha256:ddf52008bce1be455fe2b22d780b6693259aaf97b16383b6372f4b22dd33ad66 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy source code
COPY . .

RUN go build -o app .

FROM alpine:latest@sha256:8a1f59ffb675680d47db6337b49d22281a139e9d709335b492be023728e11715

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE ${API_PORT:-8080}

CMD ["./app"]
