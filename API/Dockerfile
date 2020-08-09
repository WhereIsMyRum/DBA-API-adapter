FROM golang:latest
WORKDIR /go/src
RUN apt-get update
RUN apt install -y nodejs 
RUN apt install -y npm
RUN npm install -g nodemon
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./go.mod .

WORKDIR /go/src/cmd/web/
RUN go mod download
RUN go build