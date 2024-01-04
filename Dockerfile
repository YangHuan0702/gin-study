FROM golang:1.20.12

WORKDIR /halosky

COPY . .

RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go build

EXPOSE 8080


CMD ["./gin-study"]

 
