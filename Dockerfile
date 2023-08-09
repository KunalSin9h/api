# Build Stage
FROM golang:1.20-alpine AS builder

WORKDIR /api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd ./cmd/

RUN CGO_ENABLED=0 go build -o api ./cmd/api/*.go

# Run Stage

FROM alpine

WORKDIR /api

COPY assets ./assets/
COPY --from=builder /api/api .

ENV HOST 0.0.0.0
ENV PORT 9999

EXPOSE ${PORT}

ENTRYPOINT [ "./api" ]