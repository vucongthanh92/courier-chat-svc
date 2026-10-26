FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS build

ARG TARGETOS
ARG TARGETARCH

RUN apk add --no-cache ca-certificates git tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o chat-service .

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=build /app/chat-service /app/chat-service
COPY --from=build /app/config /app/config

EXPOSE 5002 6002 7002 8002

USER nonroot:nonroot
ENTRYPOINT ["/app/chat-service"]
