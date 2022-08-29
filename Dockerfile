FROM golang:1.18

WORKDIR /docker-go

ENV GO111MODULE=on

# Note: A better solution would be to copy an existing go.mod into the image

COPY . ./

RUN apt-get update && apt-get install -y \
    zip \
    unzip \
    nano

# re-compile
ENV CGO_ENABLED=0
# RUN go get github.com/githubnemo/CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN chmod +x /docker-go

ENTRYPOINT CompileDaemon -build="go build -v -o /docker-go/app" -command="./app"