FROM golang:latest AS builder

LABEL maintainer="gusta"
LABEL env="production"

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o server main.go

EXPOSE 8080

ENTRYPOINT ["./server"]
#Final stage

# FROM scratch

# COPY  --from=builder /app/server /server



