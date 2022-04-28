FROM golang:latest

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app

EXPOSE 8080

# mkdir gqlgen-todos
# cd gqlgen-todos
# go mod init github.com/ono-hiroshi/gqlgen-todos
# go get github.com/99designs/gqlgen
# go run github.com/99designs/gqlgen init
# ./server.go
# ファイル自動生成
# go run -mod=mod github.com/99designs/gqlgen generate
# go run server.go