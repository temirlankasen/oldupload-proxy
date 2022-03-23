FROM golang:1.17 AS build

WORKDIR /usr/local/go/src/proxy

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /server


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 80

USER nonroot:nonroot

ENTRYPOINT ["/server"]

