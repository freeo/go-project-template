FROM golang:1.16.4-alpine3.13 AS builder
##  application source files
# Compile stage
ADD . /dockerdev
WORKDIR /dockerdev
RUN go build -o /main

# Final stage
FROM alpine:3.13

EXPOSE 8081
WORKDIR /
COPY --from=builder /server /
CMD ["/main"]