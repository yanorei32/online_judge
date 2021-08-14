# Develop
FROM golang:1.16.5 as dev-env

ENV PATH=/root/go/bin:$PATH

WORKDIR "/app"
RUN go get -v github.com/rubenv/sql-migrate/...

EXPOSE 1323

# Build
FROM golang:1.16.5 as build-env

WORKDIR /work
COPY . .
RUN go mod download
RUN go build

# Product
FROM gcr.io/distroless/base-debian10 as prod-env

COPY --from=build-env /work/backend /

EXPOSE 1323

ENTRYPOINT [ "/backend" ]
