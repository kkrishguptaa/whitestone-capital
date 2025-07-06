FROM golang:1.24.4-alpine@sha256:68932fa6d4d4059845c8f40ad7e654e626f3ebd3706eef7846f319293ab5cb7a AS builder

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
