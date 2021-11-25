# Develop
FROM golang:1.17.3 as dev-env

ENV PATH=/root/go/bin:$PATH

WORKDIR "/app"
RUN go get -v github.com/cosmtrek/air

EXPOSE 1323

# Build
FROM golang:1.17.3 as build-env

WORKDIR /work
COPY . .
RUN go mod download
RUN go build

# Product
FROM gcr.io/distroless/base-debian10 as prod-env

COPY --from=build-env /work/judge /

EXPOSE 1323

ENTRYPOINT [ "/judge" ]
