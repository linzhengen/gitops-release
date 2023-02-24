FROM golang:1.20 as build

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app /src

FROM gcr.io/distroless/static
WORKDIR /

COPY --from=build /app /app

ENTRYPOINT ["/app"]