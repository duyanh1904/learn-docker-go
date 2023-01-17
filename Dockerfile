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

RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)" -- \
    -t https://github.com/denysdovhan/spaceship-prompt \
    -a 'SPACESHIP_PROMPT_ADD_NEWLINE="false"' \
    -a 'SPACESHIP_PROMPT_SEPARATE_LINE="false"' \
    -p git \
    -p ssh-agent \
    -p https://github.com/zsh-users/zsh-autosuggestions \
    -p https://github.com/zsh-users/zsh-completions

# re-compile
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN chmod +x /docker-go

ENTRYPOINT CompileDaemon -build="go build -v -o /docker-go/app" -command="./app"