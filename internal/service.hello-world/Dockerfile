FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build/internal/service.hello-world
RUN go build -o /build/service.hello-world 

FROM alpine
COPY --from=builder /build/service.hello-world  /app/
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/service.hello-world"]