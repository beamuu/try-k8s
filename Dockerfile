FROM golang:1.19-alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./out/myapp ./app/cmd/main.go


FROM alpine:3.9
RUN apk add ca-certificates
COPY --from=build /app/out/myapp /run/myapp

EXPOSE 8080

ENTRYPOINT ["/run/myapp"]