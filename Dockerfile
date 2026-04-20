# syntax=docker/dockerfile:1

FROM golang:1.25-bookworm AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG VERSION=unknown
RUN GIN_MODE=release CGO_ENABLED=0 GOOS=linux go build -trimpath \
	-ldflags "-s -w -X main.buildString=${VERSION}" \
	-o /out/hnrss .

FROM gcr.io/distroless/static-debian12:nonroot
ARG CONTAINER_NAME=hnrss
LABEL org.opencontainers.image.title="${CONTAINER_NAME}"
COPY --from=builder /out/hnrss /hnrss
USER nonroot:nonroot
EXPOSE 9000
ENTRYPOINT ["/hnrss"]
CMD ["-bind", "0.0.0.0:9000"]
