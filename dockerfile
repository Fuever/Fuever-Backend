FROM golang:alpine
WORKDIR /root/
COPY . ./
ENV GOPROXY="https://goproxy.io"
ENV FUEVER_DB=fuever_db
ENV FUEVER_GO=fuever_go
ENV FUEVER_CACHE=fuever_cache
RUN go mod download && go build ./cmd/boot.go
RUN chmod 777 boot
EXPOSE 8080