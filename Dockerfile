FROM golang:1.7.4
RUN mkdir -p /go/src/github.com/kidsdynamic/childrenlab_v2
ADD . /go/src/github.com/kidsdynamic/childrenlab_v2/
WORKDIR /go/src/github.com/kidsdynamic/childrenlab_v2
RUN go build -o main .
CMD ["/go/src/github.com/kidsdynamic/childrenlab_v2/main"]

EXPOSE 8111
