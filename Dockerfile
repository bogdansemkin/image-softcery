FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go get github.com/lib/pq
RUN go build -o image-softcery ./cmd/main.go

CMD ["./image-softcery"]