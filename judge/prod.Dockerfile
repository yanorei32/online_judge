FROM golang:1.16.5 as build-env
WORKDIR /work
COPY . .
RUN go mod download
RUN go build -o judge

FROM gcr.io/distroless/cc-debian10
COPY --from=build-env /work/judge /

ENTRYPOINT [ "/judge" ]
