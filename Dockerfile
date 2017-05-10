FROM golang:1.8.1
MAINTAINER LittleRobot daisukeayanami@gmail.com

# Install beego & bee
RUN go get -u github.com/astaxie/beego
RUN go get -u github.com/beego/bee

ADD / /go/

CMD ["bee", "run"]
