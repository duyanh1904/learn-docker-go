FROM golang:1.18

WORKDIR /docker-go

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY all file go ./

COPY ./app/main.go ./
COPY . /docker-go
# To initialize a project with go module, create go.mod
RUN go mod init shipping-go

# Add missing and/or remove unused modules
RUN go mod tidy

# This will bring all the vendors to your projects /vendor directory so that 
# you don't need to get the modules again if working from another machine on this project.
RUN go mod vendor

# RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /docker-go/app .
RUN pwd

# CMD ["./app"]