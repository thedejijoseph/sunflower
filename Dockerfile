
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./sunflower ./main.go


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/sunflower .
EXPOSE 8080
ENTRYPOINT ["./sunflower"]
