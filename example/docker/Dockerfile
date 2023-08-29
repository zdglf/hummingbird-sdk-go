ARG BASE=golang:1.17-alpine
FROM ${BASE} AS builder


# set the working directory
WORKDIR /simple-driver

COPY . .

RUN go build -ldflags "-s -w" -o  ./example/cmd/driver-linux-amd64  ./example/cmd

FROM alpine:3.12

COPY --from=builder /simple-driver/example/cmd/driver-linux-amd64 /app/

RUN mkdir /logs

STOPSIGNAL SIGKILL

# Start
ENTRYPOINT [ "/app/driver-linux-amd64" ]

#docker buildx build --platform linux/amd64 -t registry.cn-shanghai.aliyuncs.com/winc-link/simple-driver:1.0 -f example/docker/Dockerfile . --push

