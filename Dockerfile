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

# build the executable
# -o allows to specify name or location of executable. 
# Here it is built and given the name hn-scraper
RUN go build -o hn-scraper



FROM template

# copy executable
COPY --from=template hn-scraper/hn-scraper /hn-scraper


ENTRYPOINT ["./hn-scraper"]
