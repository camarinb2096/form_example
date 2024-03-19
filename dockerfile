FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o form-api cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/form-api /app/form-api

EXPOSE 8070

CMD ["./form-api"]


