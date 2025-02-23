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

RUN apk add --no-cache git
RUN git clone --depth 1 --branch v0.25.0 ${SOURCE_REPO}/prometheus/alertmanager.git /alertmanager
WORKDIR /alertmanager/

RUN go mod edit -go 1.20 \
    && go get -u golang.org/x/net@v0.17.0

RUN go build -ldflags="-s -w" -o alertmanager cmd/alertmanager/main.go && \
    chown -R 64535:64535 /alertmanager/ && \
    chmod 0700 /alertmanager/alertmanager

FROM $BASE_DISTROLESS

COPY --from=artifact /alertmanager/alertmanager /bin/

ENTRYPOINT ["/bin/alertmanager"]
