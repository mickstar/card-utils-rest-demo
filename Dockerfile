FROM golang:alpine as builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /main .
RUN sha256sum /main

FROM alpine:latest as deploy
WORKDIR /
COPY --from=builder /main ./main
RUN sha256sum main_exec
EXPOSE 8080
ENTRYPOINT ["./main"]