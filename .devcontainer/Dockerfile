FROM golang:1.21.5-bullseye

ENV GO111MODULE on

WORKDIR /workspace

RUN go mod init sample

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install golang.org/x/tools/gopls@latest