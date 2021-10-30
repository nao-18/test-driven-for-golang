FROM golang:1.16-alpine

WORKDIR /go/src/github.com/nao-18/test-driven-for-golang
COPY ./src .

# RUN mkdir /usr/local/go/src/tmp

RUN apk upgrade --update && \
  apk --no-cache add git

# add package to build.
RUN apk add build-base
RUN apk add alpine-sdk

RUN go get -u github.com/cosmtrek/air && \
  # go get -u github.com/onsi/ginkgo/ginkgo && \
  # go get -u github.com/onsi/gomega && \
  go build -o /go/bin/air github.com/cosmtrek/air

# RUN touch $HOME/.ack-ginkgo-rc
# RUN echo "ACK_GINKGO_RC=true" >> .ack-ginkgo-rc

CMD ["air", "-c", ".air.toml"]

EXPOSE 9000