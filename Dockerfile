FROM golang:1.21.3-bullseye

RUN go version
ENV GOPATH=/

COPY ./ ./

#install postgres
RUN apt-get update
RUN apt-get -y install postgresql-client

#make wait-for-postgres.sh
RUN chmod +x wait-for-postgres.sh

#build go app
RUN go mod download
RUN go build -o todo-app ./cmd/main.go