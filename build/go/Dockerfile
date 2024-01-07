FROM golang:1.20-alpine

ENV CGO_ENABLED 1
ENV GOPATH /go

WORKDIR /go/src/app
RUN apk add sqlite gcc musl-dev

COPY app/go.mod app/go.sum ./
RUN go mod download
RUN go install github.com/99designs/gqlgen@latest && \
    go install github.com/cosmtrek/air@latest && \
    go install entgo.io/ent/cmd/ent@latest

ENV GOOS linux
ENV GOARCH amd64

ENV JwtSecret="123456789"

EXPOSE 8080

CMD go mod tidy && air