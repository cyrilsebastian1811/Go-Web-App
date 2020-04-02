FROM golang:1.14.1 AS builder
WORKDIR /go/src/github.com/suhas1602/gowebapp/
COPY . .
RUN cd main && go get -d -v
RUN cd main && CGO_ENABLED=0 GOOS=linux go build -a

FROM alpine as final
WORKDIR /
RUN apk add --update tzdata
COPY --from=builder /go/src/github.com/suhas1602/gowebapp/main/main .
CMD ["./main"]