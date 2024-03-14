FROM golang:latest

WORKDIR /app


COPY go.mod go.sum


RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /form-api

EXPOSE 8070

CMD ["/form-api"]


