FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /app/bin/app .

FROM gcr.io/distroless/base-debian11 as release-debian
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]
