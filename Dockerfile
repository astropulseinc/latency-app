# Build the manager binary
FROM golang:1.21 AS builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace
COPY go.mod go.mod
# COPY go.sum go.sum
RUN go mod download

COPY main.go main.go
COPY index.html index.html

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o latency main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/latency .
COPY --from=builder /workspace/index.html .
USER 65532:65532

ENTRYPOINT ["/latency"]