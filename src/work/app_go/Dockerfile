FROM golang:1.13
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get github.com/gin-gonic/gin
CMD ["go", "run", "gin.go"]
