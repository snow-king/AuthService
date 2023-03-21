FROM golang:1.18-alpine
RUN mkdir app
WORKDIR /app
COPY . ./
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init
RUN go mod download
RUN  go build -o /opt/cmd ./

FROM alpine:latest
WORKDIR /opt/
COPY .env ./
COPY --from=0 /opt/cmd ./cmd
EXPOSE 8080
CMD ["./cmd"]
