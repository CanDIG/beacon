FROM golang:1.13 AS build
WORKDIR /go/src
COPY go ./go
COPY vendor ./vendor
COPY main.go go.mod go.sum ./

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]
