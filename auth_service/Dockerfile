FROM golang:alpine AS build

WORKDIR /app
COPY . .

RUN go mod download && go mod verify
RUN go build -v -o /app/bin/app .

FROM gcr.io/distroless/base-debian11 as release-debian

WORKDIR /
COPY --from=build /app/bin/app /app
COPY /rbac/auth_model.conf/ \
     /auth_service/policy.csv \
     /auth_service/blacklist.txt \
     ./
     
RUN chmod 600 auth_model.conf
EXPOSE 8000
USER nonroot:nonroot
ENTRYPOINT ["/app"]
