FROM golang:1.9
RUN mkdir -p /go/src/github.com/kidsdynamic/childrenlab_v2/app
ADD ./app/build /go/src/github.com/kidsdynamic/childrenlab_v2/app/
WORKDIR /go/src/github.com/kidsdynamic/childrenlab_v2/app
CMD ["/go/src/github.com/kidsdynamic/childrenlab_v2/app/main"]

EXPOSE 8111
