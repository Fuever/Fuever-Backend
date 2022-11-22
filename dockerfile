FROM golang:alpine
WORKDIR /root/
COPY . ./
ENV GOPROXY="https://goproxy.io"
RUN go mod download && go build ./cmd/boot.go
RUN chmod 777 boot
EXPOSE 8080
