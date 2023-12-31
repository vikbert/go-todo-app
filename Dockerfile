FROM golang:1.12.7-alpine3.10 AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY . .
## all routing packages
RUN go get github.com/vikbert/go-todo-app
RUN go get github.com/gorilla/mux
## all routing packages
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./main.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/test --port 8080

