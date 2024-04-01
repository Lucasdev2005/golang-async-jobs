### building project. ###
    FROM golang:1.21.7 as builder
    ARG PATH_DIR
    WORKDIR /app
    COPY . .
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build ./internal/$PATH_DIR
### end. ###

### Get binary file and running on separate container. ###
    FROM golang:1.21.7-alpine3.18 as binary
    WORKDIR /app
    COPY --from=builder /app/build .
    CMD ["./build"]
### end. ### 