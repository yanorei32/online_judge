# Develop
FROM golang:1.16.5 as dev-env

WORKDIR "/app"

EXPOSE 1323

# Build
FROM golang:1.16.5 as build-env

WORKDIR /work
COPY . .
RUN go mod download
RUN go build

# Product
FROM gcr.io/distroless/base-debian10 as prod-env

COPY --from=build-env /work/judge /

EXPOSE 1323

ENTRYPOINT [ "/judge" ]
