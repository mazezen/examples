FROM golang:1.23.8


RUN mkdir -p /letterpress

COPY go.mod go.sum /letterpress/

WORKDIR /letterpress

RUN go mod download
RUN go mod tidy

COPY . /letterpress

RUN go build -o /usr/bin/letterpress /letterpress/cmd/api

EXPOSE 8080 8080
ENTRYPOINT [ "/usr/bin/letterpress" ]