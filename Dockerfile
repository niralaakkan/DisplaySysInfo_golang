FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD displaylocalipv2.go /app
WORKDIR /app
RUN go build -o ipaddr displaylocalipv2.go
CMD ["/app/ipaddr"]