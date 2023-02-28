FROM golang:latest

WORKDIR /build

COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /server

ENTRYPOINT [ "/server" ]