FROM golang:1.21-alpine
LABEL authors="1"

RUN go version
ENV GOPATH=/

COPY .. ./

RUN go mod download
RUN go build -o My house ./main.go

CMD [".Home/My house"]
ENTRYPOINT ["top", "-b"]
