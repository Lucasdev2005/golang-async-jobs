### building project. ###
    FROM golang:1.21.7 as builder
    ARG PATH_DIR
    WORKDIR /app
    COPY . .
    RUN go build -o build ./internal/$PATH_DIR
### end. ###

### Get binary file and running on separate container. ###
    FROM golang:1.21.7 as binary
    COPY --from=builder /app/build .
    CMD ["./build"]
### end. ###