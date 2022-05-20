FROM golang:1.18

WORKDIR /app

COPY ./ /app

RUN apt-get install git

RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon@latest

ENTRYPOINT CompileDaemon --build="go build -o=./out/webapp" --command=./out/webapp
