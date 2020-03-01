FROM golang:alpine AS build-env
WORKDIR $GOPATH/src/github.com/ironcore864/go-hello-http
COPY . .
RUN go build -o hello

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/ironcore864/go-hello-http/hello /app/
ENTRYPOINT app
EXPOSE 8080/tcp
