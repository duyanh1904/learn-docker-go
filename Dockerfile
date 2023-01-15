FROM golang:1.18

WORKDIR /docker-go

ENV GO111MODULE=on
ENV CGO_ENABLED=1

# Note: A better solution would be to copy an existing go.mod into the image

COPY . ./

RUN apt-get update && apt-get install -y \
    zip \
    unzip \
    nano

#install swag cli
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Run the following command to install the Go protocol buffers plugin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# gen go code
#RUN protoc -I=. --go_out=. ./pb/grpc-server/v1/address.proto

# re-compile
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN chmod +x /docker-go

ENTRYPOINT CompileDaemon -build="go build -v -o /docker-go/app" -command="./app"