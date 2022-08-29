FROM golang:1.18

WORKDIR /docker-go

# Note: A better solution would be to copy an existing go.mod into the image
RUN go mod init shipping-go
COPY ./app/main.go ./
# re-compile
ENV CGO_ENABLED=0
RUN go get github.com/githubnemo/CompileDaemon

# Determine required modules and download them
RUN go mod tidy
RUN go build -v -o /docker-go/app

ENTRYPOINT CompileDaemon -build="go build -v -o /docker-go/app" -command="./app"

RUN chmod +x /docker-go/app
