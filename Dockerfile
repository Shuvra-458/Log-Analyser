FROM golang:1.25.4-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=Linux go build -o golog
FROM alpine:latest 
WORKDIR /app
COPY --from=builder /app/golog .
RUN adduser -D gologuser
USER gologuser
ENTRYPOINT [ "./golog" ]
