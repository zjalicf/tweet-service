FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /app .

FROM gcr.io/distroless/base-debian11 as release-debian
WORKDIR /
COPY --from=build /app /app
EXPOSE 8000
USER nonroot:nonroot
CMD ["/app"]
