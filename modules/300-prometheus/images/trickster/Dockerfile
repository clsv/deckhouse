ARG BASE_DISTROLESS
ARG BASE_GOLANG_20_ALPINE

FROM $BASE_GOLANG_20_ALPINE as artifact

ARG GOPROXY
ARG SOURCE_REPO

ENV GOPROXY=${GOPROXY} \
    SOURCE_REPO=${SOURCE_REPO} \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache make git
RUN git clone --depth 1 --branch v1.1.5 ${SOURCE_REPO}/trickstercache/trickster.git /trickster
WORKDIR /trickster/

RUN go mod edit -go 1.20 \
    && go get -u golang.org/x/net@v0.17.0 \
    && go get -u github.com/prometheus/client_golang@v1.17.0
RUN make build && \
    chown -R 64535:64535 /trickster/ && \
    chmod 0700 /trickster/OPATH/trickster

FROM $BASE_GOLANG_20_ALPINE as entrypoint

COPY /entrypoint/ /app/
WORKDIR /app/

ARG GOPROXY

ENV GOPROXY=${GOPROXY} \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -ldflags="-s -w" -o entrypoint && \
    chown -R 64535:64535 /app/ && \
    chmod 0700 /app/entrypoint

FROM $BASE_GOLANG_20_ALPINE as artifact1

COPY /entrypoint/ /app/
WORKDIR /app/

ARG GOPROXY
ENV GOPROXY=${GOPROXY}

ENV GOPROXY=${GOPROXY} \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -ldflags="-s -w" -o entrypoint && \
    chown -R 64535:64535 /app/ && \
    chmod 0700 /app/entrypoint

FROM $BASE_DISTROLESS

COPY --from=artifact /trickster/OPATH/trickster /usr/local/bin/
COPY --from=entrypoint /app/entrypoint /usr/local/bin/

ENTRYPOINT ["entrypoint"]
