FROM docker.io/library/golang:1.19-bullseye as build
WORKDIR /go/src
COPY . /go/src/
ENV CGO_ENABLED=0
RUN go mod tidy && \
    go build -o desafio-clean-architecture github.com/rodrigoafernandes/desafio-clean-architecture/cmd/ordersystem

FROM scratch
WORKDIR /usr/local/bin
COPY --from=build /go/src/desafio-clean-architecture /usr/local/bin/desafio-clean-architecture
COPY --from=build /go/src/cmd/ordersystem/.env /usr/local/bin/.env
ENTRYPOINT ["./desafio-clean-architecture"]