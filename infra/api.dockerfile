# syntax=docker/dockerfile:1

# Build stage
# ---
FROM golang:1.21-alpine as build

RUN go install github.com/go-task/task/v3/cmd/task@latest

WORKDIR /app

COPY go.mod go.sum ./
COPY Taskfile.yaml .

RUN go mod download

COPY cmd/api cmd/api
COPY internal internal

RUN --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} task build.api

# Run stage
# ---
FROM alpine

COPY --from=build /app/build/api .

CMD ./api