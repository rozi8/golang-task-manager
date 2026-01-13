# contoh dockerfdocile single stage 

# FROM golang:1.22
# WORKDIR /home/rozi12/golang-task-manager/cmd/server/main.go
# COPY . .
# RUN go build -o server ./cmd/server
# CMD ["./server"]

#contoh dockerfile multi stage

#stage 1 build 
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./cmd/server

#stage 2 run
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
