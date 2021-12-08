FROM arm64v8/golang:1.17

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

#CMD ["test/go/scan_server"]