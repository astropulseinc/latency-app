FROM golang:1.27.3 AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src
COPY go.mod main.go index.html ./
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -trimpath -ldflags="-s -w" -o latency .

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /src/latency /
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/latency"]
