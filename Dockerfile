# Keep this in line with the go version in go.mod file
ARG GO_VERSION=1.27.4

# ==========================================
# Setup: Shared Toolchain
# ==========================================
FROM golang:${GO_VERSION}-alpine AS base
WORKDIR /app
ENV CGO_ENABLED=0

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# ==========================================
# Local Development Configuration
# ==========================================
FROM base AS local

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    go install github.com/air-verse/air@v1.67.4 

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

# ==========================================
# Build Stage
# ==========================================
FROM base AS build

ARG VERSION
ARG COMMIT
ARG TARGETOS
ARG TARGETARCH

COPY . .

RUN mkdir -p /out
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath \
      -ldflags="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}" \
      -o /out/app ./cmd/api 

# ==========================================
# Production Build
# ==========================================

FROM gcr.io/distroless/static-debian12:nonroot AS prod

COPY --from=build /out/app /app 

EXPOSE 8080

USER nonroot:nonroot 

ENTRYPOINT ["/app"]


