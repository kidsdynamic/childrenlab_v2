FROM golang:1.8.3
RUN mkdir -p /go/src/github.com/kidsdynamic/childrenlab_v2
ADD ./app/build /go/src/github.com/kidsdynamic/childrenlab_v2/app/
WORKDIR /go/src/github.com/kidsdynamic/childrenlab_v2
CMD ["/go/src/github.com/kidsdynamic/childrenlab_v2/app/main"]

EXPOSE 8111
