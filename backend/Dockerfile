# ref: https://github.com/ddymko/Go-Beanstalk
# https://ddymko.medium.com/deploying-docker-go-to-aws-beanstalk-e91207fe21d6
# https://walkthinksleep.tistory.com/entry/AWS-Elastic-Beanstalk-Docker-Go
FROM golang:alpine AS builder

WORKDIR /build
COPY . /build
RUN go build -o coldchain-api

EXPOSE 80
ENTRYPOINT ["./coldchain-api"]
