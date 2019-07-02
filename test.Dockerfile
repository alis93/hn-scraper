FROM golang:latest as template

# To help with go modules
ENV GO111MODULE=on

RUN mkdir /hn-scraper
WORKDIR /hn-scraper

# for dependencies
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENTRYPOINT ["go","test","./..."]
