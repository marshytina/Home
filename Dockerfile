FROM golang:latest
LABEL authors="1"

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o main .

CMD ["./main"]
ENTRYPOINT ["top", "-b"]
