FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build/internal/service.addition
RUN go build -o /build/service.addition 

FROM alpine
COPY --from=builder /build/service.addition  /app/
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/service.addition"]