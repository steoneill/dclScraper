FROM golang:1.16-alpine

WORKDIR $GOPATH/src/github.com/steoneill/dclScraper

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["dclScraper"]