FROM golang:rc-alpine3.15 as build
ENV GO111MODULE=on GOPROXY=direct
RUN mkdir -p /home/app/
WORKDIR /home/app
COPY . .
RUN go mod download -x
RUN go build -o main docker-build/main.go

FROM alpine:latest
RUN mkdir -p /home/app/
WORKDIR /home/app/
COPY --from=builder /home/app .
CMD printenv && ./main -host 0.0.0.0 -port 9090