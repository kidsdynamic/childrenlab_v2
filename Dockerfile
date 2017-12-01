FROM golang:1.9
RUN mkdir -p /go/src/github.com/kidsdynamic/childrenlab_v2/templates
ADD ./build /go/src/github.com/kidsdynamic/childrenlab_v2/
ADD ./templates /go/src/github.com/kidsdynamic/childrenlab_v2/templates
WORKDIR /go/src/github.com/kidsdynamic/childrenlab_v2
CMD ["/go/src/github.com/kidsdynamic/childrenlab_v2/main"]

EXPOSE 8111
