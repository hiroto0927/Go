FROM golang:1.21.5-bullseye

ENV GO111MODULE on

WORKDIR /workspace

RUN go mod init sample

RUN go install github.com/cosmtrek/air@latest \
    && go install github.com/ramya-rao-a/go-outline@latest \
    && go install golang.org/x/tools/gopls@latest

RUN git config --global --add safe.directory /workspace 