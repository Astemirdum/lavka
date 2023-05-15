FROM golang:1.20-alpine AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
RUN apk update --no-cache && apk add --no-cache tzdata
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download && go mod verify
COPY . .
RUN go build -ldflags="-s -w" -o /app/lavka ./cmd/main.go

FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/lavka /app/lavka
COPY --from=builder /build/.env /app

CMD ["./lavka"]

