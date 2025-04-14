ARG version=alpine
FROM golang:${version} AS builder
WORKDIR /app
COPY app.go .
RUN go build app.go


FROM scratch
COPY --from=builder /app/app .
ENV NAME="Jane Doe"
ENV CHARACTER="SpiderMan"

CMD ["./app"]