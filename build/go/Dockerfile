FROM golang:1.20-alpine

ENV GOPATH /go
ENV CGO_ENABLED 0
ENV GO111MODULE on

WORKDIR /go/src/app

COPY app/go.mod app/go.sum ./
RUN go mod download
RUN go install github.com/99designs/gqlgen@latest && \
    go install github.com/cosmtrek/air@latest && \
    go install github.com/volatiletech/sqlboiler/v4@v4.14.0  && \
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest && \
    go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN apk add sqlite gcc musl-dev
ENV CGO_ENABLED 1

EXPOSE 8080

CMD go mod tidy && air