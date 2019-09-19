FROM golang:1.13 AS build
WORKDIR /src/beacon
ENV CGO_ENABLED=0
COPY . .
RUN go build

FROM scratch
ENV GIN_MODE=release
COPY --from=build /src/beacon /
EXPOSE 8080/tcp
ENTRYPOINT ["/beacon"]
