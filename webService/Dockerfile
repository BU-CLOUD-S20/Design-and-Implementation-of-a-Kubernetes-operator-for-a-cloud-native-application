
FROM golang:1.12.7 

WORKDIR /Users/yuanwei/Desktop/projDemo

COPY . /Users/yuanwei/Desktop/projDemo
RUN go build -o /Users/yuanwei/Desktop/projDemo/k8s-microservice
EXPOSE 1323/tcp
ENTRYPOINT ["/Users/yuanwei/Desktop/projDemo/k8s-microservice"]


