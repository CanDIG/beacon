FROM golang:1.13 AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN go build -o beacon

FROM scratch
ENV GIN_MODE=release
COPY --from=build /src/beacon /
EXPOSE 8080/tcp
ENTRYPOINT ["/beacon"]
env BEACON_VERBOSE=
ENV CANDIG_URL=http://candig-server
