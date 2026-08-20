#STAGE 1
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /app/bin/server .


#STAGE 2
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /app/bin/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
