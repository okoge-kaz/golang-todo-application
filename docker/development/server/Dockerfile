FROM golang:1.19

WORKDIR /go/server
COPY ./server .

RUN go mod download
RUN go mod tidy

CMD ["air", "-c", "server/.air.toml"]
