FROM golang:1.16.5

ENV GOPATH=
ENV PATH=/root/go/bin:$PATH

COPY go.* ./
RUN go mod download
RUN go get -v github.com/rubenv/sql-migrate/...

WORKDIR "/app"

ENTRYPOINT ["go", "run", "main.go"]
