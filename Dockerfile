FROM golang:latest AS builder

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -buildvcs=false -o /build/omsms-server-init .

FROM gcr.io/distroless/base-debian12
LABEL authors="october1234"

COPY --from=builder /build/omsms-server-init /omsms-server-init
ENTRYPOINT ["/omsms-server-init"]
