FROM golang:1.22
WORKDIR /home/rozi12/golang-task-manager/cmd/server/main.go
COPY . .
RUN go build -o server ./cmd/server
CMD ["./server"]

