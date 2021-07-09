FROM golang:1.16.5

ENV GOPATH=

COPY go.* ./
RUN go mod download

WORKDIR "/app"

ENTRYPOINT ["go", "run", "main.go"]
