FROM golang:1.8.5-jessie

# install dep
RUN go get github.com/golang/dep/cmd/dep
# install gin
RUN go get github.com/codegangsta/gin

WORKDIR /go/src/app

ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock

RUN dep ensure --vendor-only

ADD src src

CMD ["go", "run", "src/main.go"]