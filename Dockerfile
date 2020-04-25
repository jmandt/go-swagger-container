FROM golang:1.13-alpine

RUN apk update git

WORKDIR /go/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY cmd/blueprint ./

COPY ./config /config

EXPOSE 8080

# ENTRYPOINT ["/{ARG_BIN}"]

# Used in Makefile - `make container` step