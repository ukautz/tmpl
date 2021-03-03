ARG GO_VERSION="1.16"
FROM golang:${GO_VERSION} AS build

WORKDIR /src
ADD . .
ENV CGO_ENABLED=0
RUN make test
RUN make build

FROM scratch

COPY --from=build /src/dist/tmpl /usr/bin/tmpl

ENTRYPOINT [ "/usr/bin/tmpl" ]


