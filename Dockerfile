FROM golang:1.18

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o=./out/webapp" --command=./out/webapp
